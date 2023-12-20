package day12

import (
	utils "advent-of-code/utils"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sync"
	// "math"
	"strings"
	"time"
)

func lineValid(line string, grouping []int) bool {
	blocks := []int{}
	for _, block := range strings.Split(line, ".") {
		if block != "" {
			blocks = append(blocks, len(block))
		}
	}
	if len(blocks) != len(grouping) {
		return false
	}
	for i, block := range blocks {
		if block != grouping[i] {
			return false
		}
	}
	return true
}


func soFarValid(line string, grouping []int) bool {
	line_so_far := strings.Split(line, "?")[0]
	blocks_so_far := []int{}
	for _, block := range strings.Split(line_so_far, ".") {
		if block != "" {
			blocks_so_far = append(blocks_so_far, len(block))
		}
	}
	blocks_min := []int{}
	for _, block := range strings.Split(strings.Replace(line, "?", "", -1), ".") {
		if block != "" {
			blocks_min = append(blocks_min, len(block))
		}
	}
	blocks_max := []int{}
	line_max := line
	for strings.Contains(line_max, "?") {
		index_first_question_mark := strings.Index(line_max, "?")
		replace_with := ""
		if index_first_question_mark > 0 {
			if line_max[index_first_question_mark-1] == '.' {
				replace_with = string(line_max[index_first_question_mark-1])
			} else {
				replace_with = "."
			}
		} else {
			if line_max[index_first_question_mark+1] == '.' {
				replace_with = "#"
			} else {
				replace_with = "."
			}
		}

		line_max = line_max[:index_first_question_mark] + replace_with + line_max[index_first_question_mark+1:]
	}
	// line_max := strings.Replace(line,    "????????????????", "#.#.#.#.#.#.#.#.#", -1)
	// line_max = strings.Replace(line_max, "???????????????",  "#.#.#.#.#.#.#.#.", -1)
	// line_max = strings.Replace(line_max, "??????????????",   "#.#.#.#.#.#.#.#", -1)
	// line_max = strings.Replace(line_max, "?????????????",    "#.#.#.#.#.#.#.", -1)
	// line_max = strings.Replace(line_max, "????????????",     "#.#.#.#.#.#.#", -1)
	// line_max = strings.Replace(line_max, "???????????",      "#.#.#.#.#.#.", -1)
	// line_max = strings.Replace(line_max, "??????????",       "#.#.#.#.#.#", -1)
	// line_max = strings.Replace(line_max, "?????????",        "#.#.#.#.#.", -1)
	// line_max = strings.Replace(line_max, "????????",         "#.#.#.#.#", -1)
	// line_max = strings.Replace(line_max, "???????",          "#.#.#.#.", -1)
	// line_max = strings.Replace(line_max, "??????",           "#.#.#.#", -1)
	// line_max = strings.Replace(line_max, "?????",            "#.#.#.", -1)
	// line_max = strings.Replace(line_max, "????",             "#.#.#", -1)
	// line_max = strings.Replace(line_max, "???",              "#.#.", -1)
	// line_max = strings.Replace(line_max, "??",               "#.#", -1)
	// line_max = strings.Replace(line_max, "?",                "#.", -1)
	for _, block := range strings.Split(line_max, ".") {
		if block != "" {
			blocks_max = append(blocks_max, len(block))
		}
	}
	if len(blocks_max) < len(grouping) {
		// println("too few blocks")
		return false
	}
	if len(blocks_min) > len(grouping) {
		// println("too many blocks")
		return false
	}
	for i, block := range blocks_so_far {
		if i < (len(blocks_so_far)-1) && (block != grouping[i]) {
			return false
		}
		if i == (len(blocks_so_far)-1) && (block > grouping[i]) {
			return false
		}
	}

	remaining_block_length := 0
	for i := len(blocks_so_far); i < len(grouping); i++ {
		remaining_block_length += grouping[i]
	}
	if (remaining_block_length + len(grouping) - len(blocks_so_far) - 1) > (len(line) - len(line_so_far)) {
		// println("too little of line left")
		return false
	}
	return true
}

var SYMBOLS []rune

func countValidCombinationsLineRec(line []rune, grouping []int, index int) int {
	if index == len(line) {
		// all question marks have been replaced, check line
		if lineValid(string(line), grouping) {
			return 1
		} else {
			return 0
		}	
	}
	if !soFarValid(string(line), grouping) {
		return 0
	}
	if line[index] != rune('?') {
		return countValidCombinationsLineRec(line, grouping, index+1)
	} else {
		count := 0
		for _, symbol := range SYMBOLS {
			line[index] = symbol
			count += countValidCombinationsLineRec(line, grouping, index+1)
			line[index] = '?'
		}
		return count
	}
}

var wg sync.WaitGroup


func Part1_concurrent(lines [][]rune, groupings [][]int) int {
	solution := 0
	ch := make(chan int, len(lines))
	for i := 0; i < len(lines); i++ {
		wg.Add(1)
		i := i
		go func() {
			defer wg.Done()
			ch <- countValidCombinationsLineRec(lines[i], groupings[i], 0)
		}()
	}
	wg.Wait()
	close(ch)
	for i := range ch {
		solution += i
	}
	return solution
}

func Part1_sequential(lines [][]rune, groupings [][]int) int {
	solution := 0
	for i := 0; i < len(lines); i++ {
		solution += countValidCombinationsLineRec(lines[i], groupings[i], 0)
	}
	return solution
}

func Solve(runAs string) {

	// SOLUTION_I := 0
	// SOLUTION_II := 0

	SYMBOLS = []rune{'.', '#'}

	dir, err := os.Getwd()
	utils.CheckError(err)

	var inputFile = fmt.Sprintf("%s/day12/input_%s.txt", dir, runAs)

	println(inputFile)

	readFile, err := os.Open(inputFile)
	utils.CheckError(err)
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	lines := [][]rune{}
	groupings := [][]int{}
	for fileScanner.Scan() {
		split_line := strings.Split(fileScanner.Text(), " ")
		line := []rune(split_line[0])
		lines = append(lines, line)
		grouping_str := strings.Split(split_line[1], ",")
		grouping_int := make([]int, len(grouping_str))
		for i := 0; i < len(grouping_str); i++ {
			grouping_int[i], err = strconv.Atoi(grouping_str[i])
			utils.CheckError(err)
		}
		groupings = append(groupings, grouping_int)
	}

	// fmt.Println(lines)
	// fmt.Println(groupings)
	// fmt.Println(SYMBOLS)

	// time the two runs
	start := time.Now()
	SOLUTION_I := Part1_concurrent(lines, groupings)
	elapsed := time.Since(start)
	SOLUTION_I = Part1_sequential(lines, groupings)
	elapsed2 := time.Since(start)
	fmt.Println("Concurrent:", elapsed)
	fmt.Println("Sequential:", elapsed2)
	println("The solution for part I is:", SOLUTION_I)

	// modify input lines for part II
	lines_unfolded := [][]rune{}
	groupings_unfolded := [][]int{}
	for i := 0; i < len(lines); i++ {
		line_unfolded := []rune{}
		grouping_unfolded := []int{}
		for j := 0; j < 5; j++ {
			line_unfolded = append(line_unfolded, lines[i]...)
			grouping_unfolded = append(grouping_unfolded, groupings[i]...)
			if j < 4 {
				line_unfolded = append(line_unfolded, '?')			}
		}
		lines_unfolded = append(lines_unfolded, line_unfolded)
		groupings_unfolded = append(groupings_unfolded, grouping_unfolded)
	}

	// fmt.Println(lines_unfolded)
	// fmt.Println(groupings_unfolded)
	start = time.Now()
	SOLUTION_II := Part1_concurrent(lines_unfolded, groupings_unfolded)
	elapsed = time.Since(start)
	fmt.Println("Concurrent:", elapsed)
	println("The solution for part II is:", SOLUTION_II)
}
