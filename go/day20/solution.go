package day20

import (
	utils "advent-of-code/utils"
	"bufio"
	"fmt"
	"os"

	// "math"
	"strings"
)

func b2i(b bool) uint8 {
	if b {
		return 1
	}
	return 0
}

// https://siongui.github.io/2017/06/03/go-find-lcm-by-gcd/
// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
			t := b
			b = a % b
			a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
			result = LCM(result, integers[i])
	}

	return result
}

type Pulse uint8

const (
	Low  Pulse = 0
	High Pulse = 1
)

type module interface {
	get_receivers() []string
	process(Pulse, string) ([]string, Pulse)
}

type broadcaster struct {
	receivers []string
}

func (b *broadcaster) process(pulse Pulse, _ string) ([]string, Pulse) {
	return b.receivers, pulse
}
func (b broadcaster) get_receivers() []string {
	return b.receivers
}

type flipflop struct {
	state     bool
	receivers []string
}
func (f *flipflop) process(pulse Pulse, _ string) ([]string, Pulse) {
	if pulse == High {
		return []string{}, pulse
	}

	f.state = !f.state

	return f.receivers, Pulse(b2i(f.state))
}
func (f flipflop) get_receivers() []string {
	return f.receivers
}

type conjunction struct {
	memory    map[string]Pulse
	receivers []string
}

func (c *conjunction) process(pulse Pulse, name string) ([]string, Pulse) {
	c.memory[name] = pulse

	for _, value := range c.memory {
		if value == Low {
			return c.receivers, High
		}
	}

	return c.receivers, Low
}
func (c conjunction) get_receivers() []string {
	return c.receivers
}

type queued_pulse struct {
	pulse       Pulse
	from_module string
	to_module   string
}

var MODULES map[string]module

func Part1() int {

	n_low_pulses := 0
	n_high_pulses := 0

	for i := 0; i < 1000; i++ {

		pulse_queue := []queued_pulse{{Low, "button", "broadcaster"}}

		for len(pulse_queue) > 0 {

			new_pulse_queue := []queued_pulse{}

			for _, p := range pulse_queue {
				// count pulses
				if p.pulse == Low {
					n_low_pulses++
				} else {
					n_high_pulses++
				}
					
				if MODULES[p.to_module] == nil {
					continue
				}
				
				receivers, pulse := MODULES[p.to_module].process(p.pulse, p.from_module)

				for _, receiver := range receivers {
					new_pulse_queue = append(new_pulse_queue, queued_pulse{pulse, p.to_module, receiver})
				}
			}

			pulse_queue = new_pulse_queue
		}
	}
	
	return n_low_pulses * n_high_pulses
}

func Part2() int {
	parent := findParents("rx")[0]
	
	cycles := make(map[string]int)
	if lv_conj, ok := MODULES[parent].(*conjunction); ok {
		for name := range lv_conj.memory {
			cycles[name] = 0
		}
	}
	
	n_iter := 0
	rx_triggered := false
	all_cycles_detected := false

	for !rx_triggered && !all_cycles_detected && n_iter < 1_000_000 {
		n_iter++

		pulse_queue := []queued_pulse{{Low, "button", "broadcaster"}}

		for len(pulse_queue) > 0 {

			new_pulse_queue := []queued_pulse{}

			for _, p := range pulse_queue {

				if p.pulse == Low && p.to_module == "rx" {
					return n_iter
				}

				if lv_conj, ok := MODULES[p.to_module].(*conjunction); ok && p.to_module == parent {
					for name, m := range lv_conj.memory {
						if m == High && cycles[name] == 0 {
							cycles[name] = n_iter
						}
					}
				}

				if MODULES[p.to_module] == nil {
					continue
				}
				
				receivers, pulse := MODULES[p.to_module].process(p.pulse, p.from_module)

				for _, receiver := range receivers {
					new_pulse_queue = append(new_pulse_queue, queued_pulse{pulse, p.to_module, receiver})
				}
			}

			pulse_queue = new_pulse_queue

			all_cycles_detected = true
			for _, c := range cycles {
				if c == 0 {
					all_cycles_detected = false
				}
			}
		}

	}

	cycles_values := []int{}
	for _, v := range cycles {
		cycles_values = append(cycles_values, v)
	}

	return LCM(cycles_values[0], cycles_values[1], cycles_values...)
}

func readInput(inputFile string) {

	readFile, err := os.Open(inputFile)
	utils.CheckError(err)
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	MODULES = make(map[string]module)
	conjunctions := []string{}

	for fileScanner.Scan() {
		line := fileScanner.Text()
		split := strings.Split(line, " -> ")
		name := split[0]
		receivers := strings.Split(split[1], ", ")

		switch {
		case name == "broadcaster":
				MODULES[name] = &broadcaster{receivers: receivers}
		case []rune(name)[0] == '%':
				MODULES[strings.TrimPrefix(name, "%")] = &flipflop{state: false, receivers: receivers}
		case []rune(name)[0] == '&':
				conj := conjunction{memory: make(map[string]Pulse), receivers: receivers}
				strippedName := strings.TrimPrefix(name, "&")
				MODULES[strippedName] = &conj
				conjunctions = append(conjunctions, strippedName)
		}
	}

	for name, module := range MODULES {
		for _, receiver := range module.get_receivers() {
			for _, conj := range conjunctions {
				if conj == receiver {
					if conjunction, ok := MODULES[conj].(*conjunction); ok {
						conjunction.memory[name] = Low
					}
				}
			}
		}
	}
}

func findParents(module string) []string {
	parents := []string{}
	for name, m := range MODULES {
		for _, r := range m.get_receivers() {
			if r == module {
				parents = append(parents, name)
			}
		}
	}
	
	return parents
}

func Solve(runAs string) {
	dir, err := os.Getwd()
	utils.CheckError(err)

	var inputFile = fmt.Sprintf("%s/day20/input_%s.txt", dir, runAs)

	readInput(inputFile)
	SOLUTION_I := Part1()
	println("The solution for part I is:", SOLUTION_I)

	readInput(inputFile)
	SOLUTION_II := Part2()
	println("The solution for part II is:", SOLUTION_II)
}
