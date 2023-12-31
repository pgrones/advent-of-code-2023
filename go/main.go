package main

import (
	day2 "advent-of-code/day2"
	day4 "advent-of-code/day4"
	day6 "advent-of-code/day6"
	day8 "advent-of-code/day8"
	// day10 "advent-of-code/day10"
	day12 "advent-of-code/day12"
	day14 "advent-of-code/day14"
	day16 "advent-of-code/day16"
	day18 "advent-of-code/day18"
	// day20 "advent-of-code/day20"
	// day22 "advent-of-code/day22"
	// day24 "advent-of-code/day24"
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
		case "8":
			day8.Solve(os.Args[2])
		// case "10":
		// 	day10.Solve(os.Args[2])
		case "12":
			day12.Solve(os.Args[2])
		case "14":
			day14.Solve(os.Args[2])
		case "16":
			day16.Solve(os.Args[2])
		case "18":
			day18.Solve(os.Args[2])
		// case "20":
		// 	day20.Solve(os.Args[2])
		// case "22":
		// 	day22.Solve(os.Args[2])
		// case "24":
		// 	day24.Solve(os.Args[2])
	}
}