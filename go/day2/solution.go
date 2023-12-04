package day2

import (
	"bufio"
	"fmt"
	"os"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func Solve(runAs string) {
	dir, err := os.Getwd()
    check(err)

	var inputFile = fmt.Sprintf("%s/day2/input_%s.txt", dir, runAs)
	
	readFile, err := os.Open(inputFile)
  
    check(err)

    fileScanner := bufio.NewScanner(readFile)
 
    fileScanner.Split(bufio.ScanLines)
  
    for fileScanner.Scan() {
        fmt.Println(fileScanner.Text())
    }
  
    readFile.Close()
}