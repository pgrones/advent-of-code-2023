package day6

import (
	utils "advent-of-code/utils"
	"bufio"
	"fmt"
	"os"
	"strings"

	// "math"
	// "os"
	"strconv"
)

var TIMES = []int{}
var DISTANCES = []int{}

func Solve(runAs string) {
	dir, err := os.Getwd()
	utils.CheckError(err)

	var inputFile = fmt.Sprintf("%s/day6/input_%s.txt", dir, runAs)

	println(inputFile)

	readFile, err := os.Open(inputFile)
    utils.CheckError(err)
    fileScanner := bufio.NewScanner(readFile)
    fileScanner.Split(bufio.ScanLines)
  
    SCORE_I := 1
    SCORE_II := 0
	
	fileScanner.Scan()
	times_str := strings.TrimPrefix(fileScanner.Text(), "Time:")
	fileScanner.Scan()
	distances_str := strings.TrimPrefix(fileScanner.Text(), "Distance:")
	
	times_str_list := strings.Fields(times_str)
	distances_str_list := strings.Fields(distances_str)

	for _, time_str := range times_str_list {
		time_int, err := strconv.Atoi(time_str)
		utils.CheckError(err)
		TIMES = append(TIMES, time_int)	
	}
		
	for _, distance_str := range distances_str_list {
		distance_int, err := strconv.Atoi(distance_str)
		utils.CheckError(err)
		DISTANCES = append(DISTANCES, distance_int)
	}

	fmt.Println(TIMES)
	fmt.Println(DISTANCES)

	for i := 0; i < len(TIMES); i++ {

		count := 0

		time := TIMES[i]
		distance := DISTANCES[i]

		for j := 0; j < time; j++ {
			race_time := time - j
			race_distance := race_time * j
			if race_distance > distance {
				count += 1
			} else if count > 0 {
				break
			}
		}
		
		SCORE_I *= count
	}

	long_time_str := strings.Join(times_str_list, "")
	long_distance_str := strings.Join(distances_str_list, "")

	long_time, err := strconv.Atoi(long_time_str)
	utils.CheckError(err)
	long_distance, err := strconv.Atoi(long_distance_str)
	utils.CheckError(err)

	for i := 0; i < long_time; i++ {
		race_time := long_time - i
		race_distance := race_time * i
		if race_distance > long_distance {
			SCORE_II += 1
		} else if SCORE_II > 0 {
			break
		}
	}
	
	println("The solution for part I is:", SCORE_I)
	println("The solution for part II is:", SCORE_II)
}