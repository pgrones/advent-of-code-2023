package day2

import (
	utils "advent-of-code/utils"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

var CUBE_LIMITS map[string]int

var MAX_CUBES map[string]int

func checkGame(gameLine string) (int, bool) {

    gamePossible := true

    gameLine = strings.TrimLeft(gameLine, "Game ")

    split := strings.Split(gameLine, ": ")
    gameID, err := strconv.Atoi(split[0])
    
    utils.CheckError(err)
    gameLine = split[1]    
    utils.CheckError(err)

    rounds := strings.Split(gameLine, "; ")

    for i := 0; i < len(rounds); i++ {
        perColor := strings.Split(rounds[i], ", ")

        for j := 0; j < len(perColor); j++{
            split := strings.Split(perColor[j], " ")
            amount, err := strconv.Atoi(split[0])
            utils.CheckError(err)
            color := split[1]

            if amount > CUBE_LIMITS[color] {
                gamePossible = false
            }
        }
    }

    return gameID, gamePossible
}

func powerGame(gameLine string) int {

    MAX_CUBES = make(map[string]int)
    MAX_CUBES["red"] = 0
    MAX_CUBES["green"] = 0
    MAX_CUBES["blue"] = 0

    gameLine = strings.TrimLeft(gameLine, "Game ")

    split := strings.Split(gameLine, ": ")
    _, err := strconv.Atoi(split[0])
    
    utils.CheckError(err)
    gameLine = split[1]    
    utils.CheckError(err)

    rounds := strings.Split(gameLine, "; ")

    for i := 0; i < len(rounds); i++ {
        perColor := strings.Split(rounds[i], ", ")

        for j := 0; j < len(perColor); j++{
            split := strings.Split(perColor[j], " ")
            amount, err := strconv.Atoi(split[0])
            utils.CheckError(err)
            color := split[1]

            if amount > MAX_CUBES[color] {
                MAX_CUBES[color] = amount
            }
        }
    }
    return MAX_CUBES["red"] * MAX_CUBES["green"] * MAX_CUBES["blue"]
}


func Solve(runAs string) {
	dir, err := os.Getwd()
    utils.CheckError(err)

    var inputFile = fmt.Sprintf("%s/day2/input_%s.txt", dir, runAs)

    // ###### PART 1 ######

    SCORE_I := 0
    SCORE_II := 0

    // === Part I ===
    CUBE_LIMITS = make(map[string]int)
    CUBE_LIMITS["red"] = 12
    CUBE_LIMITS["green"] = 13
    CUBE_LIMITS["blue"] = 14
    println("Red cube limit:", CUBE_LIMITS["red"])
    println("Green cube limit:", CUBE_LIMITS["green"])
    println("Blue cube limit:", CUBE_LIMITS["blue"])

	

	readFile, err := os.Open(inputFile)
    utils.CheckError(err)
    fileScanner := bufio.NewScanner(readFile)
    fileScanner.Split(bufio.ScanLines)
  
    for fileScanner.Scan() {
        gameID, possible := checkGame(fileScanner.Text())
        // println(gameID, possible)
        if possible {
            SCORE_I += gameID
        }
        SCORE_II += powerGame(fileScanner.Text())
    }

    readFile.Close()

    println("The magical solutions for PART I and II as obtained by the almighty programming language GO are:")
    time.Sleep(time.Second)
    println("...")
    time.Sleep(time.Second)
    println("drum roll")
    time.Sleep(time.Second)
    println("...")
    time.Sleep(time.Second)
    println(SCORE_I, SCORE_II)
}