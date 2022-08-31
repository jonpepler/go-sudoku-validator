package main

import (
	"fmt"
	"strconv"
	"strings"
)

var sample_sudokus = [][][]int{
	{
		{8, 2, 9, 3, 6, 5, 1, 4, 7},
		{6, 4, 3, 7, 1, 8, 5, 1, 9},
		{7, 5, 1, 4, 9, 2, 6, 3, 8},
		{3, 1, 8, 5, 2, 7, 4, 9, 6},
		{5, 9, 6, 1, 4, 1, 7, 8, 2},
		{4, 7, 2, 9, 8, 6, 3, 1, 5},
		{9, 8, 4, 6, 5, 1, 2, 7, 3},
		{2, 6, 7, 8, 3, 4, 9, 5, 1},
		{9, 3, 5, 2, 7, 9, 8, 6, 4},
	},
	{
		{5, 8, 6, 4, 3, 7, 1, 9, 2},
		{1, 9, 4, 5, 8, 2, 3, 6, 7},
		{7, 2, 3, 9, 6, 1, 4, 5, 8},
		{2, 4, 7, 1, 9, 8, 6, 3, 5},
		{8, 3, 9, 6, 2, 4, 7, 2, 1},
		{6, 5, 1, 7, 2, 3, 8, 4, 9},
		{9, 7, 5, 3, 1, 6, 7, 8, 4},
		{3, 1, 8, 2, 4, 5, 9, 7, 6},
		{4, 6, 2, 8, 7, 9, 5, 1, 3},
	},
	{
		{1, 8, 3, 2, 7, 4, 6, 5, 9},
		{9, 7, 4, 5, 8, 6, 3, 2, 1},
		{2, 6, 5, 1, 9, 3, 7, 4, 8},
		{5, 9, 2, 8, 3, 1, 4, 6, 7},
		{8, 4, 6, 7, 2, 5, 9, 1, 3},
		{7, 3, 1, 4, 6, 9, 2, 8, 5},
		{3, 5, 9, 6, 4, 8, 1, 7, 2},
		{6, 1, 7, 3, 5, 2, 8, 9, 4},
		{4, 2, 8, 9, 1, 7, 5, 3, 6},
	},
}

type sudoku_point struct {
	value int
	valid bool
}

func main() {
	fmt.Println("Sample 1")
	validate_sudoku(sample_sudokus[0])
	fmt.Println("\n\nSample 2")
	validate_sudoku(sample_sudokus[1])
	fmt.Println("\n\nSample 3")
	validate_sudoku(sample_sudokus[2])
}

func validate_sudoku(sudoku [][]int) {
	rows := sudoku
	columns := transpose_sudoku(sudoku)
	grids := grid_sudoku(sudoku)

	valid := true
	sudoku_validation_map := make_map(sudoku)
	error_text := "Errors: "
	error_locs := []string{}

	for i, row := range rows {
		errors, valid_zone := check_zone((row))
		if !valid_zone {
			error_locs = append(error_locs, "Row "+strconv.Itoa(i+1))

			for _, idx := range errors {
				sudoku_validation_map[i][idx].valid = false
			}
			valid = false
		}
	}

	for i, column := range columns {
		errors, valid_zone := check_zone((column))
		if !valid_zone {
			error_locs = append(error_locs, "Column "+strconv.Itoa(i+1))
			for _, idx := range errors {
				sudoku_validation_map[idx][i].valid = false
			}
			valid = false
		}
	}

	for i, grid := range grids {
		errors, valid_zone := check_zone((grid))
		if !valid_zone {
			error_locs = append(error_locs, "Grid "+strconv.Itoa(i+1))
			for _, idx := range errors {
				sudoku_validation_map[idx/3+(i/3)*3][idx%3+(i%3)*3].valid = false
			}
			valid = false
		}
	}

	fmt.Println(valid)
	if !valid {
		fmt.Println()
		print_validation_map(sudoku_validation_map)
		fmt.Println()
		fmt.Println(error_text, strings.Join(error_locs, ", "))
	}
}

func print_validation_map(validation_map [][]sudoku_point) {
	colourReset := "\033[0m"
	colourRed := "\033[31m"
	for _, row := range validation_map {
		for _, p := range row {
			str := strconv.Itoa(p.value)
			if !p.valid {
				str = colourRed + str + colourReset
			}
			fmt.Print(" " + str + " ")
		}
		fmt.Print("\n")
	}
}

func make_map(sudoku [][]int) (validation_map [][]sudoku_point) {
	for i := 0; i < len(sudoku); i++ {
		validation_map_row := []sudoku_point{}
		for j := 0; j < len(sudoku[0]); j++ {
			validation_map_row = append(validation_map_row, sudoku_point{value: sudoku[i][j], valid: true})
		}
		validation_map = append(validation_map, validation_map_row)
	}
	return
}

func check_zone(zone []int) (errors []int, valid bool) {
	errors = get_duplicates(zone)
	return errors, len(errors) == 0
}

func get_duplicates(list []int) (duplicates []int) {
	freq := make(map[int]int)
	for _, v := range list {
		_, exists := freq[v]
		if exists {
			freq[v] += 1
		} else {
			freq[v] = 1
		}
	}
	for v, f := range freq {
		if f > 1 {
			for i, v2 := range list {
				if v2 == v {
					duplicates = append(duplicates, i)
				}
			}
		}
	}

	return
}

func transpose_sudoku(sudoku [][]int) [][]int {
	result := make([][]int, 9)
	for i := range result {
		result[i] = make([]int, 9)
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			result[i][j] = sudoku[j][i]
		}
	}
	return result
}

func grid_sudoku(sudoku [][]int) (grids [][]int) {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			grids = append(grids, []int{
				sudoku[i*3][j*3], sudoku[i*3][j*3+1], sudoku[i*3][j*3+2],
				sudoku[i*3+1][j*3], sudoku[i*3+1][j*3+1], sudoku[i*3+1][j*3+2],
				sudoku[i*3+2][j*3], sudoku[i*3+2][j*3+1], sudoku[i*3+2][j*3+2],
			})
		}
	}
	return
}
