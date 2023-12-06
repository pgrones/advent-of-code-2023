package day4

import (
	utils "advent-of-code/utils"
	"fmt"
	"os"
)

func Solve(runAs string) {
	dir, err := os.Getwd()
	utils.CheckError(err)

	var inputFile = fmt.Sprintf("%s/day4/input_%s.txt", dir, runAs)

	println(inputFile)
}