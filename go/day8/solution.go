package day8

import (
	utils "advent-of-code/utils"
	"bufio"
	"fmt"
	"os"
	"strconv"
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

func findCycles(nodes []*Node) ([][]int, []int) {
	chunks := [][]int{}
	remainders := []int{}

	for _, node := range nodes {
		CURR_INDEX = 0
		iter_count := 0

		node_list := []*Node{node}

		visited_slice := []string{}
		visited_map := make(map[string]bool)

		for {
			position_str := node_list[0].name + strconv.Itoa(CURR_INDEX)

			if visited_map[position_str] && string([]rune(position_str)[2]) == "Z" {
				// find index of previous visit
				cycle_index := 0
				z_indexes := []int{}
				for j, pos := range visited_slice {
					if pos == position_str {
						cycle_index = j
					}

					if cycle_index != 0 && string([]rune(pos)[2]) == "Z" {
						z_indexes = append(z_indexes, j)
					}
				}

				i := 0
				chunk := []int{}
				if len(z_indexes) > 1 {
					for i < len(z_indexes) - 1 {
						chunk = append(chunk, z_indexes[i + 1] - z_indexes[i])
						i++
					}
				}
				chunk = append(chunk,  iter_count - z_indexes[len(z_indexes) - 1])
				chunks = append(chunks, chunk)
				remainders = append(remainders, cycle_index)
				break
			}
			
			visited_slice = append(visited_slice, position_str)
			visited_map[position_str] = true
			node_list = update_nodes(nextTurn(), node_list)
			iter_count++
		}
	}

	return chunks, remainders
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
 
	SCORE_I += walk(start_nodes_I)

	cycles, remainders := findCycles(start_nodes_II)
	fmt.Println(cycles)
	fmt.Println(remainders)

	var curr_chunk_idcs [6]int

	iters := remainders

	for {
		// find smallest iter
		min_iter_idx := 0
		min_iter := iters[0]
		for i, iter := range iters {
			if iter < min_iter {
				min_iter_idx = i
				min_iter = iter
			}
		}
		// add corresponding cycle to iters
		iters[min_iter_idx]	+= cycles[min_iter_idx][curr_chunk_idcs[min_iter_idx]]
		curr_chunk_idcs[min_iter_idx]++
		if curr_chunk_idcs[min_iter_idx] >= len(cycles[min_iter_idx]) {
			curr_chunk_idcs[min_iter_idx] = 0
		}

		// check if all equal
		all_equal := true
		for _, iter := range iters {
			if iter != iters[0] {
				all_equal = false
				break
			}
		}
		if all_equal {
			break
		}
		// else, repeat
	}

	SCORE_II = iters[0]

	println("The solution for part I is:", SCORE_I)
	println("The solution for part II is:", SCORE_II)
}