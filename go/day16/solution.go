package day16

import (
	utils "advent-of-code/utils"
	"fmt"
	"os"
	"bufio"
	// "time"
	// "math"
	// "strings"
)


type Coord struct {
	x int
	y int
}

type Tile struct {
	symbol rune
	energized_directional map[string]bool
	energized bool
	coord Coord
}

type TileMap struct {
	tiles [][]Tile
	size_x int
	size_y int
}

var tile_map TileMap


func loadInput(inputFile string) {
	readFile, err := os.Open(inputFile)
	utils.CheckError(err)
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	tiles := [][]Tile{}
		
	i := 0
	for fileScanner.Scan() {
	    symbols := fileScanner.Text()
		line_tiles := []Tile{}
		for j, symbol := range symbols {
			line_tiles = append(line_tiles, Tile{symbol, make(map[string]bool), false, Coord{j, i}}) 
		}
		tiles = append(tiles, line_tiles)
		i++
	}

	tile_map = TileMap{tiles, len(tiles[0]), len(tiles)}
}


func paintTileMap() {
	for i := 0; i < tile_map.size_x; i++ {
		for j := 0; j < tile_map.size_y; j++ {
			if tile_map.tiles[i][j].energized {				
				print("\033[35m" + "#" + "\033[0m")
			} else {
				print(".")
			}
		}
		println()
	}
}


func propagate(coord Coord, direction string) {
	x, y := coord.x, coord.y
	if x < 0 || x >= tile_map.size_x || y < 0 || y >= tile_map.size_y {
		return
	}

	tile := tile_map.tiles[y][x]

	if tile.energized_directional[direction] {
		return
	}

	// time.Sleep(10 * time.Millisecond)
	// paintTileMap()

	tile_map.tiles[y][x].energized = true
	tile_map.tiles[y][x].energized_directional[direction] = true

	move_to := []string {direction}

	if tile.symbol == '/' {
		switch direction {
			case "left": move_to[0] = "down"
			case "right": move_to[0] = "up"
			case "up":  move_to[0] = "right"
			case "down":  move_to[0] = "left"
		}
	}
	
	if tile.symbol == '\\' {
		switch direction {
			case "left": move_to[0] = "up"
			case "right":  move_to[0] = "down"
			case "up":  move_to[0] = "left"
			case "down":  move_to[0] = "right"
		}
	}

	if tile.symbol == '|' {
		switch direction {
			case "left": fallthrough
			case "right": move_to = []string {"up", "down"}
		}
	}

	if tile.symbol == '-' {
		switch direction {
			case "up": fallthrough
			case "down": move_to = []string {"left", "right"}
		}
	}

	// propagate beam
	for _, direction := range move_to {
		switch direction {
			case "left": propagate(Coord{x-1, y}, direction)
			case "right": propagate(Coord{x+1, y}, direction)
			case "up": propagate(Coord{x, y-1}, direction)
			case "down": propagate(Coord{x, y+1}, direction)
		}
	}
}


func resetTileMap() {
	for i := 0; i < tile_map.size_x; i++ {
		for j := 0; j < tile_map.size_y; j++ {
			tile_map.tiles[i][j].energized = false
			tile_map.tiles[i][j].energized_directional = make(map[string]bool)
		}
	}
}


func Part1(start Coord, direction string) int {

	propagate(start, direction)

	// count energized tiles
	solution := 0
	for i := 0; i < tile_map.size_x; i++ {
		for j := 0; j < tile_map.size_y; j++ {
			if tile_map.tiles[i][j].energized {				
				solution++
			}
		}
	}

	resetTileMap()
	return solution
}

func Part2() int {
	solution := 0
	tmp := 0
	opt_coord := Coord{0, 0}
	for i := 0; i < tile_map.size_x; i++ {
		tmp = Part1(Coord{i, 0}, "down")
		if tmp > solution {
			solution = tmp
			opt_coord = Coord{i, 0}
		}
		tmp = Part1(Coord{i, tile_map.size_y-1}, "up")
		if tmp > solution {
			solution = tmp
			opt_coord = Coord{i, tile_map.size_y-1}
		}
	}
	for j := 0; j < tile_map.size_y; j++ {
		tmp = Part1(Coord{0, j}, "right")
		if tmp > solution {
			solution = tmp
			opt_coord = Coord{0, j}
		}
		tmp = Part1(Coord{tile_map.size_x-1, j}, "left")
		if tmp > solution {
			solution = tmp
			opt_coord = Coord{tile_map.size_x-1, j}
		}
	}
	fmt.Println("Optimal starting coordinate:", opt_coord)
	return solution
}

func Solve(runAs string) {
	dir, err := os.Getwd()
	utils.CheckError(err)

	var inputFile = fmt.Sprintf("%s/day16/input_%s.txt", dir, runAs)
	loadInput(inputFile)
	
    SOLUTION_I := Part1(Coord{0, 0}, "right")
    SOLUTION_II := Part2()

	println("The solution for part I is:", SOLUTION_I)
	println("The solution for part II is:", SOLUTION_II)
}