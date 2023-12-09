package day8

import (
	utils "advent-of-code/utils"
	"bufio"
	"fmt"
	"os"
	"strings"
	// "math"
)

var INSTRUCTION []rune
var CURR_INDEX int


func nextTurn() string {
	turn := INSTRUCTION[CURR_INDEX]
	
	if CURR_INDEX >= len(INSTRUCTION) - 1 {
		CURR_INDEX = 0
	} else {
		CURR_INDEX++
	}

	return string(turn)
}


type Node struct {
	name string

	left_name string
	right_name string

	left *Node
	right *Node
}


func walk(nodes []*Node) int {
	all_Z := true

	turn := nextTurn()
	
	for i := range nodes {	
		if string([]rune(nodes[i].name)[2]) != "Z" {
			all_Z = false
		}

		if len(nodes) == 1 && nodes[i].name == "ZZZ" {
			return 0
		}
		
		if turn == "L" {
			nodes[i] = nodes[i].left
		} else {
			nodes[i] = nodes[i].right
		}
	}

	if all_Z {
		return 0
	} else {
		return 1 + walk(nodes)
	}
}

func update_nodes(turn string, nodes []*Node) []*Node {
	for i := range nodes {			
		if turn == "L" {
			nodes[i] = nodes[i].left
		} else {
			nodes[i] = nodes[i].right
		}
	}

	return nodes
}

func walk_rollout(nodes []*Node) int {
	iter_count := 0
	var all_Z bool
	
	for {
		if iter_count % 10000000 == 0 {
			println(iter_count)
		}
		all_Z = true
		
		for i := range nodes {
			if string([]rune(nodes[i].name)[2]) != "Z" {
				all_Z = false
			}
		}

		if all_Z {
			break
		}
			
		nodes = update_nodes(nextTurn(), nodes)
		iter_count++
	}

	return iter_count
}


func Solve(runAs string) {
	dir, err := os.Getwd()
	utils.CheckError(err)

	var inputFile = fmt.Sprintf("%s/day8/input_%s.txt", dir, runAs)

	println(inputFile)

	readFile, err := os.Open(inputFile)
    utils.CheckError(err)
    fileScanner := bufio.NewScanner(readFile)
    fileScanner.Split(bufio.ScanLines)
  
    SCORE_I := 0
    SCORE_II := 0

	GRAPH := []Node{}

	fileScanner.Scan()
	instruction_str := fileScanner.Text()
	INSTRUCTION = []rune(instruction_str)
	fileScanner.Scan()

    for fileScanner.Scan() {
		line := fileScanner.Text()
		split := strings.Split(line, " = ")
		node_name := split[0]
		left_right := strings.Split(strings.Trim(split[1], "()"), ", ")
		GRAPH = append(GRAPH, Node{name: node_name, left_name: left_right[0], right_name: left_right[1]})
	}

	start_nodes_I := []*Node{}
	start_nodes_II := []*Node{}

	for i := range GRAPH {
		if GRAPH[i].name == "AAA" {
			start_nodes_I = append(start_nodes_I, &GRAPH[i])
		}

		if string([]rune(GRAPH[i].name)[2]) == "A" {
			start_nodes_II = append(start_nodes_II, &GRAPH[i])
		}

		for j := range GRAPH {
			if GRAPH[j].name == GRAPH[i].left_name {
				GRAPH[i].left = &GRAPH[j]
			}

			if GRAPH[j].name == GRAPH[i].right_name {
				GRAPH[i].right = &GRAPH[j]
			}
		}
	}
 
	// SCORE_I += walk(start_nodes_I)

	CURR_INDEX = 0
	SCORE_II += walk_rollout(start_nodes_II)

	println("The solution for part I is:", SCORE_I)
	println("The solution for part II is:", SCORE_II)
}