package main

import (
	day2 "advent-of-code/day2"
	day4 "advent-of-code/day4"
	day6 "advent-of-code/day6"
	"os"
)

func main() {

	switch os.Args[1]{
		case "2":
			day2.Solve(os.Args[2])
		case "4":
			day4.Solve(os.Args[2])
		case "6":
			day6.Solve(os.Args[2])
	}
}