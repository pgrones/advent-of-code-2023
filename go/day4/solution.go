package day4

import (
	utils "advent-of-code/utils"
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

var ALL_MATCHES = []int{}

func remove[T any](slice []T, s int) []T {
    return append(slice[:s], slice[s+1:]...)
}

func calc_matches(card string) int {

	split := strings.Split(card, " | ")
	winning := strings.Split(split[0], " ")
	having := strings.Split(split[1], " ")

	for i, entri := range having {
		if entri == "" || entri == " " {
			having = remove(having, i)
		}
	}
	for i, entri := range winning {
		if entri == "" || entri == " " {
			winning = remove(winning, i)
		}
	}

	matches := 0

	for _, entri := range having {
		for _, entrj := range winning {
			if entri == entrj {
				matches += 1
				break;
			}
		}
	}

	return matches
}


func accumulate_cards(id int, matches int) int {
	n_cards := 1
	for i := id + 1; i <= id + matches; i++ {
		n_cards += accumulate_cards(i, ALL_MATCHES[i])
	}
	return n_cards
}


func Solve(runAs string) {
	dir, err := os.Getwd()
	utils.CheckError(err)

	var inputFile = fmt.Sprintf("%s/day4/input_%s.txt", dir, runAs)

	println(inputFile)

	readFile, err := os.Open(inputFile)
    utils.CheckError(err)
    fileScanner := bufio.NewScanner(readFile)
    fileScanner.Split(bufio.ScanLines)
  
    SCORE_I := 0
    SCORE_II := 0
	
	card := ""
	matches := 0

    for fileScanner.Scan() {
        // check number of matches for current card
		card = fileScanner.Text()
		card = strings.Split(card, ": ")[1]
		matches = calc_matches(card)
		// calculate value of card
		ALL_MATCHES = append(ALL_MATCHES, matches)
		SCORE_I += int(math.Pow(float64(2), float64(matches - 1)))
    }

	for i, matches := range ALL_MATCHES {
		SCORE_II += accumulate_cards(i, matches)
	}

	println("The solution for part I is:", SCORE_I)
	println("The solution for part II is:", SCORE_II)
}