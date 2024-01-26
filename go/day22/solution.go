package day22

import (
	utils "advent-of-code/utils"
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
	"sort"
)


type block struct {
	start []int
	end []int
	x []int
	y []int
	z []int
	supported_by_blocks []int
}


func readInput(inputFile string) []block {

	blocks := []block{}

	readFile, err := os.Open(inputFile)
	utils.CheckError(err)
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		start_end := strings.Split(fileScanner.Text(), "~")
		start_strings := strings.Split(start_end[0], ",")
		end_strings := strings.Split(start_end[1], ",")

		var start = []int{}
		var end = []int{}

		for i := 0; i < 3; i++ {
			
			start_int, err := strconv.Atoi(start_strings[i])
			utils.CheckError(err)
			
			end_int, err := strconv.Atoi(end_strings[i])
			utils.CheckError(err)
			
			start = append(start, int(start_int))
			end = append(end, int(end_int))
		}
		b := block {
			start: start, end: end,
			x: []int{start[0], end[0]},
			y: []int{start[1], end[1]},
			z: []int{start[2], end[2]},
			supported_by_blocks: []int{}}
		blocks = append(blocks, b)
	}
	return blocks
}


func orderCoords(bs []block) []block {

	blocks := deepCopyBlocks(bs)

	for _, b := range blocks {
		for i := 0; i < 3; i++ {
			if b.start[i] > b.end[i] {
				start_buf := b.start[i]
				b.start[i] = b.end[i]
				b.end[i] = start_buf
			}
			switch i {
				case 0: {
					b.x[0] = b.start[i] // 5 // = []int{b.start[0], b.end[0]}
					b.x[1] = b.end[i]
				}
				case 1: {
					b.y[0] = b.start[i]
					b.y[1] = b.end[i]
				}
				case 2: {
					b.z[0] = b.start[i]
					b.z[1] = b.end[i]
				}
			}
		}
		b.y = []int{b.start[1], b.end[1]}
		b.z = []int{b.start[2], b.end[2]}	
	}
	return blocks
}


func sortBlocks_by_z(blocks []block) {
	sort.Slice(blocks, func(i, j int) bool {
		return blocks[i].z[0] < blocks[j].z[0]
	})
}


func deepCopyBlocks(blocks []block) []block {

	blocks_copy := []block{}

	for _, b := range blocks {
		block_copy := b

		block_copy.start = append([]int{}, b.start...)
		block_copy.end = append([]int{}, b.end...)
		block_copy.x = append([]int{}, b.x...)
		block_copy.y = append([]int{}, b.y...)
		block_copy.z = append([]int{}, b.z...)
		block_copy.supported_by_blocks = append([]int{}, b.supported_by_blocks...)

		blocks_copy = append(blocks_copy, block_copy)
	}

	return blocks_copy
}


func applyGravity(bs []block) []block {

	blocks := deepCopyBlocks(bs)
	
	for i := range blocks {
		// find block below
		z_new := 0
		for j := 0; j < i; j++ {
			if !(blocks[j].x[0] > blocks[i].x[1] || blocks[j].x[1] < blocks[i].x[0]) && !(blocks[j].y[0] > blocks[i].y[1] || blocks[j].y[1] < blocks[i].y[0]) {
				if blocks[j].z[1] > z_new {
					z_new = blocks[j].z[1]
					blocks[i].supported_by_blocks = []int{j}
				} else if blocks[j].z[1] == z_new {
						blocks[i].supported_by_blocks = append(blocks[i].supported_by_blocks, j)
				}
			}
		}

		fall_dist := blocks[i].z[0] - z_new - 1

		// update z of block
		blocks[i].z[0] -= fall_dist
		blocks[i].z[1] -= fall_dist
		blocks[i].start[2] -= fall_dist
		blocks[i].end[2] -= fall_dist
	}

	return blocks
}


func Part1(blocks []block) int {

	solution := 0

	for i := range blocks {
		// check if this block is the only support of any other block
		safe := true
		for j := range blocks {
			if len(blocks[j].supported_by_blocks) == 1 && blocks[j].supported_by_blocks[0] == i {
				safe = false
			}
		}
		if safe {
			solution++
		}
	}

	return solution
}


func findSupportedBlocks(blocks []block, supported_by map[int]bool) []int {
	supported_blocks := []int{}
	for i, block := range blocks {
		supports_i := true
		for _, support := range block.supported_by_blocks {
			if !supported_by[support] {
				supports_i = false
				break
			}
		}
		if supports_i && len(block.supported_by_blocks) > 0 {
			supported_blocks = append(supported_blocks, i)
		}
	}
	return supported_blocks
}


func Part2(blocks []block) int {

	solution := 0

	for i := range blocks {
		// println("======== Block", i, "=========")

		falling_blocks := make(map[int]bool)
		falling_blocks[i] = true
		supported_blocks := findSupportedBlocks(blocks, falling_blocks)

		for len(supported_blocks) > 0 {
			// fmt.Println(falling_blocks)
			// fmt.Println(supported_blocks)
			new_supported_blocks := []int{}
			for _, b := range findSupportedBlocks(blocks, falling_blocks) {
				if !falling_blocks[b] {
					falling_blocks[b] = true
					new_supported_blocks = append(new_supported_blocks, b)
				}
			}
			supported_blocks = new_supported_blocks
		}

		solution += len(falling_blocks) - 1
	}

	return solution
}


func Solve(runAs string) {
	dir, err := os.Getwd()
	utils.CheckError(err)

	var inputFile = fmt.Sprintf("%s/day22/input_%s.txt", dir, runAs)
	println(inputFile)

	blocks := readInput(inputFile)
	blocks = orderCoords(blocks)
	sortBlocks_by_z(blocks)

	// let blocks fall
	blocks = applyGravity(blocks)

    SOLUTION_I := Part1(blocks)
	println("The solution for part I is:", SOLUTION_I)

    SOLUTION_II := Part2(blocks)
	println("The solution for part II is:", SOLUTION_II)	
}