package day0

import (
	utils "advent-of-code/utils"
	"fmt"
	"os"
	// "math"
	// "os"
	// "strings"
)


func Solve(runAs string) {
	dir, err := os.Getwd()
	utils.CheckError(err)

	var inputFile = fmt.Sprintf("%s/day0/input_%s.txt", dir, runAs)

	println(inputFile)

	// readFile, err := os.Open(inputFile)
    // utils.CheckError(err)
    // fileScanner := bufio.NewScanner(readFile)
    // fileScanner.Split(bufio.ScanLines)
  
    // SCORE_I := 0
    // SCORE_II := 0
	
    // for fileScanner.Scan() {
    //     // do something
	// }

	// println("The solution for part I is:", SCORE_I)
	// println("The solution for part II is:", SCORE_II)
}