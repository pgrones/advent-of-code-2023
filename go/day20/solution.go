package day20

import (
	utils "advent-of-code/utils"
	"fmt"
	"os"
	"bufio"
	// "math"
	"strings"
)


type module interface {
	pulse_in()
	pulses_out()
}


type button struct {
	receivers []string{}
}


type broadcaster struct {

}


type flipflop struct {

}


type conjunction struct {

}


func readInput(inputFile string) {

	readFile, err := os.Open(inputFile)
	utils.CheckError(err)
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	
	for fileScanner.Scan() {
		// do something
	}
}


func Solve(runAs string) {
	dir, err := os.Getwd()
	utils.CheckError(err)

	var inputFile = fmt.Sprintf("%s/day20/input_%s.txt", dir, runAs)
	input := readInput(inputFile)

	
	SOLUTION_I := Part1(input)
	println("The solution for part I is:", SOLUTION_I)

	// SOLUTION_II := Part2()
	// println("The solution for part II is:", SOLUTION_II)
}