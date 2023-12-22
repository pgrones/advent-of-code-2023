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


func Part1(input [][]rune) int {
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
				
				if i == 9 && j == 2 {
					println(index)
				}
				input[index][j] = 'O'
				input[i][j] = '.'
			}
		}
	}

	// compute solution
	// fmt.Println(input)
	solution := 0
	for i, line := range input {
		fmt.Println(string(line))
		n_rocks := 0
		for _, symbol := range line {
			if symbol == 'O' {
				// println("here")
				n_rocks++
			}
		}
		// println(n_rocks)
		solution += n_rocks * (len(input) - i)
	}
	return solution
}


func Solve(runAs string) {
	dir, err := os.Getwd()
	utils.CheckError(err)

	var inputFile = fmt.Sprintf("%s/day14/input_%s.txt", dir, runAs)

	println(inputFile)

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

	fmt.Println(input)
	start := time.Now()
    SOLUTION_I := Part1(input)
	elapsed := time.Since(start)
	fmt.Println("Finished in:", elapsed)
    SOLUTION_II := 0

	println("The solution for part I is:", SOLUTION_I)
	println("The solution for part II is:", SOLUTION_II)
}