package day18

import (
	utils "advent-of-code/utils"
	"bufio"
	"fmt"
	"os"

	// "math"
	"strconv"
	"strings"
)


type TrenchMap struct {
	trenches [][]rune
	colors [][]string
	min_row int
	max_row int
	min_col int
	max_col int
}


func digTrench(trench_map TrenchMap, curr_row int, curr_col int, direction string, distance int, color string) (TrenchMap, int, int) {
	
	if direction == "U" {
		if curr_row - distance < trench_map.min_row {
			trench_map = extendMap(trench_map, "U", trench_map.min_row - (curr_row - distance))
		}
		// dig
		for i := 1; i <= distance; i++ {
			trench_map.trenches[curr_row - trench_map.min_row - i][curr_col - trench_map.min_col] = '#'
			trench_map.colors[curr_row - trench_map.min_row - i][curr_col - trench_map.min_col] = color
		}
		curr_row -= distance
	}

	if direction == "D" {
		if curr_row + distance > trench_map.max_row {
			trench_map = extendMap(trench_map, "D", curr_row + distance - trench_map.max_row)
		}
		// dig
		for i := 1; i <= distance; i++ {
			trench_map.trenches[curr_row - trench_map.min_row + i][curr_col - trench_map.min_col] = '#'
			trench_map.colors[curr_row - trench_map.min_row + i][curr_col - trench_map.min_col] = color
		}
		curr_row += distance
	}

	if direction == "R" {
		if curr_col + distance > trench_map.max_col {
			trench_map = extendMap(trench_map, "R", curr_col + distance- trench_map.max_col)
		}
		// dig
		for i := 1; i <= distance; i++ {
			trench_map.trenches[curr_row - trench_map.min_row][curr_col - trench_map.min_col + i] = '#'
			trench_map.colors[curr_row - trench_map.min_row][curr_col - trench_map.min_col + i] = color
		}
		curr_col += distance
	}

	if direction == "L" {
		if curr_col - distance < trench_map.min_col {
			trench_map = extendMap(trench_map, "L", trench_map.min_col - (curr_col - distance))
		}
		// dig
		for i := 1; i <= distance; i++ {
			trench_map.trenches[curr_row - trench_map.min_row][curr_col - trench_map.min_col - i] = '#'
			trench_map.colors[curr_row - trench_map.min_row][curr_col - trench_map.min_col - i] = color
		}
		curr_col -= distance
	}

	return trench_map, curr_row, curr_col
}


func extendMap(trench_map TrenchMap, direction string, distance int) (TrenchMap) {
	n_rows := (trench_map.max_row - trench_map.min_row) + 1
	n_cols := (trench_map.max_col - trench_map.min_col) + 1

	switch direction {
		case "U": {
			line := []rune{}
			color_line := []string{}
			for i := 0; i < n_cols; i++ {
				line = append(line, '.')
				color_line = append(color_line, "#ffffff")
			}
			for i := 0; i < distance; i++ {
				line_copy := make([]rune, len(line))
				copy(line_copy, line)
				trench_map.trenches = append([][]rune{line_copy}, trench_map.trenches...)

				color_line_copy := make([]string, len(color_line))
				copy(color_line_copy, color_line)
				trench_map.colors = append([][]string{color_line_copy}, trench_map.colors...)
			}
			trench_map.min_row -= distance
		}
		case "D": {
			line := []rune{}
			color_line := []string{}
			for i := 0; i < n_cols; i++ {
				line = append(line, '.')
				color_line = append(color_line, "#ffffff")
			}
			for i := 0; i < distance; i++ {
				line_copy := make([]rune, len(line))
				copy(line_copy, line)
				trench_map.trenches = append(trench_map.trenches, line_copy)

				color_line_copy := make([]string, len(color_line))
				copy(color_line_copy, color_line)
				trench_map.colors = append(trench_map.colors, color_line_copy)
			}
			trench_map.max_row += distance
		}
		case "R": {
			for i := 0; i < distance; i++ {
				for j := 0; j < n_rows; j++ {
					trench_map.trenches[j] = append(trench_map.trenches[j], '.')
					trench_map.colors[j] = append(trench_map.colors[j], "#ffffff")
				}
			}
			trench_map.max_col += distance
		}
		case "L": {
			for i := 0; i < distance; i++ {
				for j := 0; j < n_rows; j++ {
					trench_map.trenches[j] = append([]rune{'.'}, trench_map.trenches[j]...)
					trench_map.colors[j] = append([]string{"#ffffff"}, trench_map.colors[j]...)
				}
			}
			trench_map.min_col -= distance
		}
	}

	return trench_map
}

func paintMap (trench_map TrenchMap) {
	f1, _ := os.Create("./day18/output.txt")
	f2, _ := os.Create("./day18/output.html")
	defer f1.Close()
	defer f2.Close()
	f2.WriteString("<div style=\"font-family:Ubuntu Mono\">")
	for i := 0; i <= trench_map.max_row - trench_map.min_row; i++ {
		for j := 0; j <= trench_map.max_col - trench_map.min_col; j++ {
			color := trench_map.colors[i][j]
			// r, _ := strconv.ParseInt(color[1:3], 16, 0)
			// g, _ := strconv.ParseInt(color[3:5], 16, 0)
			// b, _ := strconv.ParseInt(color[5:7], 16, 0)
			// fmt.Print("\033[38;2;" + strconv.Itoa(int(r)) + ";" + strconv.Itoa(int(g)) + ";" + strconv.Itoa(int(b)) + "m" + string(trench_map.trenches[i][j]) + "\033[0m")
			f1.WriteString(string(trench_map.trenches[i][j]))
			f2.WriteString(fmt.Sprintf("<span style=\"color:%s\">%s</span>", color, string(trench_map.trenches[i][j])))
		}
		// println()
		f1.WriteString("\n")
		f2.WriteString("<br/>\n")
	}
	f2.WriteString("</div>")
}

func fillTrench(trench_map TrenchMap) TrenchMap {

	var hash_found bool
	var pointing string
	var index int
	
	for i, row := range trench_map.trenches {
		
		indices := []int{}
		hash_found = false
		pointing = ""

		for j, symbol := range row {
			if symbol == '#' && !hash_found {
				hash_found = true
				if i > 0 && trench_map.trenches[i-1][j] == '#' {
					pointing = "U"
				} else if i < len(row)-1 && trench_map.trenches[i+1][j] == '#' {
					pointing = "D"
				} else {
					pointing = ""
				}
				index = j
			}
			if hash_found && (symbol == '.' || j == len(row)-1) {
				hash_found = false
				x := j - 1
				if j == len(row) - 1 {
					x = j
				}
				if x > 0 && trench_map.trenches[i][x-1] == '#' && ((pointing == "U" && trench_map.trenches[i-1][x] == '#') || (pointing == "D" && trench_map.trenches[i+1][x] == '#')) {
					continue
				}
				indices = append(indices, index)
			}
		}
		for j := 0; j < len(indices); j += 2 {
			for i := indices[j] + 1; i < indices[j+1]; i++ {
				if row[i] == '.' {
					row[i] = '@'
				}
			}
		}
	}
	return trench_map
}


func countTrenchVolume(trench_map TrenchMap) int {
	volume := 0
	for i := 0; i < len(trench_map.trenches); i++ {
		for j := 0; j < len(trench_map.trenches[0]); j++ {
			if trench_map.trenches[i][j] == '#' || trench_map.trenches[i][j] == '@' {
				volume++
			}
			// if volume % 1000 == 0 {
			// 	println(volume)
			// }
		}
	}
	return volume
}


func Part1(directions []string, distances []int, hex_colors []string) int {

	trench_map := TrenchMap{[][]rune{{'#'}}, [][]string{{"#ffffff"}}, 0, 0, 0, 0}
	curr_col := 0
	curr_row := 0

	var direction string
	var distance int
	var color string

	for i := 0; i < len(directions); i++ {
		direction = directions[i]
		distance = distances[i]
		color = hex_colors[i]
		trench_map, curr_row, curr_col = digTrench(trench_map, curr_row, curr_col, direction, distance, color)
	}

	// paintMap(trench_map)
	// fmt.Println(trench_map.trenches)

	// fill trench
	trench_map = fillTrench(trench_map)
	
	// paint trench map
	paintMap(trench_map)

	// count trenches
	return countTrenchVolume(trench_map)
}


func Solve(runAs string) {
	dir, err := os.Getwd()
	utils.CheckError(err)

	var inputFile = fmt.Sprintf("%s/day18/input_%s.txt", dir, runAs)

	// PART I
	directions := []string{}
	distances := []int{}
	hex_colors := []string{}

	readFile, err := os.Open(inputFile)
    utils.CheckError(err)
    fileScanner := bufio.NewScanner(readFile)
    fileScanner.Split(bufio.ScanLines)
	
    for fileScanner.Scan() {
        line := fileScanner.Text()
		line_split := strings.Split(line, " ")
		directions = append(directions, line_split[0])
		distance, err := strconv.Atoi(line_split[1])
		distances = append(distances, distance)
		utils.CheckError(err)
		hex_colors = append(hex_colors, strings.Trim(line_split[2], "()"))
	}

    SOLUTION_I := Part1(directions, distances, hex_colors)
	println("The solution for part I is:", SOLUTION_I)

	// PART II
	directions = []string{}
	distances = []int{}
	hex_colors = []string{}

	readFile, err = os.Open(inputFile)
    utils.CheckError(err)
    fileScanner = bufio.NewScanner(readFile)
    fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
        line := fileScanner.Text()
		line_split := strings.Split(line, " ")
		hex_number := strings.Trim(line_split[2], "()")

		distance, _ := strconv.ParseInt(hex_number[1:6], 16, 0)
		distances = append(distances, int(distance))

		direction_int := hex_number[6]
		direction := ""
		switch direction_int {
			case '0':
				direction = "R"
			case '1':
				direction = "D"
			case '2':
				direction = "L"
			case '3':
				direction = "U"
		}
		directions = append(directions, direction)

		hex_colors = append(hex_colors, hex_number)
	}

	fmt.Println(distances)
	fmt.Println(directions)

	SOLUTION_II := Part1(directions, distances, hex_colors)
	println("The solution for part II is:", SOLUTION_II)
}