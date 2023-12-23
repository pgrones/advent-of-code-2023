package day14

import (
	utils "advent-of-code/utils"
	"fmt"
	"os"
	"bufio"
	"time"
	// "math"
	// "strings"
)


func moveNorth(input [][]rune) {
	for i := 0; i < len(input); i++ {
		if i == 0 {
			continue
		}
		for j := 0; j < len(input[i]); j++ {
			if input[i][j] == 'O' && input[i-1][j] == '.' {
				index := 0
				for k := i; k > 0; k-- {
					if input[k - 1][j] != '.' {
						index = k
						break
					}
					
				}
				input[index][j] = 'O'
				input[i][j] = '.'
			}
		}
	}
}


func computeLoad(input [][]rune) int {
	load := 0
	for i, line := range input {
		n_rocks := 0
		for _, symbol := range line {
			if symbol == 'O' {
				n_rocks++
			}
		}
		load += n_rocks * (len(input) - i)
	}
	return load
}


func Part1(input [][]rune) int {
	moveNorth(input)
	return computeLoad(input)
}


func pivotInput(input [][]rune) [][]rune {
	new_input := make([][]rune, len(input))
	for i := 0; i < len(input); i++ {
		new_input[i] = make([]rune, len(input[0]))
	}
	for i := 0; i < len(input[0]); i++ {
		for j := 0; j < len(input); j++ {
			new_input[i][len(input)-j-1] = input[j][i]
		}
	}
	return new_input
}


func Part2(input [][]rune) int {
	known_configurations := make(map[string]int)
	var (
		loop_start int
		loop_end int
	)
	for i := 0; i < 1_000_000_000; i++ {
		for j := 0; j < 4; j++ {
			moveNorth(input)
			input = pivotInput(input)
		}
		key := ""
		for _, line := range input {
			key += string(line)
		}
		if known_configurations[key] > 0 {
			println("Repetition detected at iteration:", i)
			loop_start = known_configurations[key]
			loop_end = i
			break
		} else {
			known_configurations[key] = i
		}
	}
	remaining_cycles := (1_000_000_000 - loop_start) % (loop_end - loop_start)
	for i := 0; i < remaining_cycles - 1; i++ {
		for j := 0; j < 4; j++ {
			moveNorth(input)
			input = pivotInput(input)
		}
	}
	return computeLoad(input)
}


func loadInput(inputFile string) [][]rune {
	readFile, err := os.Open(inputFile)
    utils.CheckError(err)
    fileScanner := bufio.NewScanner(readFile)
    fileScanner.Split(bufio.ScanLines)
	input := [][]rune{}
	line := ""
    for fileScanner.Scan() {
        line = fileScanner.Text()
		input = append(input, []rune(line))
	}
	return input
}


func Solve(runAs string) {
	dir, err := os.Getwd()
	utils.CheckError(err)

	var inputFile = fmt.Sprintf("%s/day14/input_%s.txt", dir, runAs)

	input := loadInput(inputFile)

	start := time.Now()
	SOLUTION_I := Part1(input)
	elapsed := time.Since(start)
	println("The solution for part I is:", SOLUTION_I)
	fmt.Println("Finished in:", elapsed)

	input = loadInput(inputFile)
	start = time.Now()
    SOLUTION_II := Part2(input)
	elapsed = time.Since(start)
	println("The solution for part II is:", SOLUTION_II)
	fmt.Println("Finished in:", elapsed)

}