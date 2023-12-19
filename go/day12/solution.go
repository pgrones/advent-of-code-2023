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
)

func lineValid(line string, grouping []int) bool {
	blocks := []string{}
	for _, block := range strings.Split(line, ".") {
		if block != "" {
			blocks = append(blocks, block)
		}
	}
	fmt.Println(blocks)
	if len(blocks) != len(grouping) {
		return false
	}
	for i, group := range grouping {
		if len(blocks[i]) != group {
			return false
		}
	}
	return true
}


var SYMBOLS []rune

func countValidCombinationsLineRec(line []rune, grouping []int, index int, ch chan int) { // int {
	// fmt.Println(string(line))
	// fmt.Println(index)
	return
	if index == len(line) {
		// all question marks have been replaced, check line
		if lineValid(string(line), grouping) {
			ch <- 1
			// return 1
		} // else {
			// return 0
		// }
		return
	}
	if line[index] != rune('?') {
		// return countValidCombinationsLineRec(line, grouping, index+1, ch)
		countValidCombinationsLineRec(line, grouping, index+1, ch)
		return
	} else {
		// count := 0
		for _, symbol := range SYMBOLS {
			line[index] = symbol
			// count += countValidCombinationsLineRec(line, grouping, index+1, ch)
			countValidCombinationsLineRec(line, grouping, index+1, ch)
			line[index] = '?'
		}
		return
		// return count
	}
}

var wg sync.WaitGroup


func Part1(lines [][]rune, groupings [][]int) int {
	solution := 0
	ch := make(chan int)
	for i := 0; i < len(lines); i++ {
		wg.Add(1)
		i := i
		go func(i int) {
			defer wg.Done()
			println("test", i)
			// countValidCombinationsLineRec(lines[i], groupings[i], 0, ch)
		}(i)
		// solution += countValidCombinationsLineRec(lines[i], groupings[i], 0, ch)
	}
	wg.Wait()
	for i := range ch {
		solution += i
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

	SOLUTION_I := Part1(lines, groupings)
	println("The solution for part I is:", SOLUTION_I)

	// SOLUTION_I := Part2(lines, groupings)
	// println("The solution for part II is:", SOLUTION_II)
}
