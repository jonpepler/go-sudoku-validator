package main

import (
	"fmt"
	"reflect"
	"sort"
)

// var sudoku = [][]int{
// 	{8, 2, 9, 3, 6, 5, 1, 4, 7},
// 	{6, 4, 3, 7, 1, 8, 5, 1, 9},
// 	{7, 5, 1, 4, 9, 2, 6, 3, 8},
// 	{3, 1, 8, 5, 2, 7, 4, 9, 6},
// 	{5, 9, 6, 1, 4, 1, 7, 8, 2},
// 	{4, 7, 2, 9, 8, 6, 3, 1, 5},
// 	{9, 8, 4, 6, 5, 1, 2, 7, 3},
// 	{2, 6, 7, 8, 3, 4, 9, 5, 1},
// 	{9, 3, 5, 2, 7, 9, 8, 6, 4},
// }

var sudoku = [][]int{
	{1, 5, 2, 4, 8, 9, 3, 7, 6},
	{7, 3, 9, 2, 5, 6, 8, 4, 1},
	{4, 6, 8, 3, 7, 1, 2, 9, 5},
	{3, 8, 7, 1, 2, 4, 6, 5, 9},
	{5, 9, 1, 7, 6, 3, 4, 2, 8},
	{2, 4, 6, 8, 9, 5, 7, 1, 3},
	{9, 1, 4, 6, 3, 7, 5, 8, 2},
	{6, 2, 5, 9, 4, 8, 1, 3, 7},
	{8, 7, 3, 5, 1, 2, 9, 6, 4},
}

func main() {
	validate_sudoku(sudoku)
}

func validate_sudoku(sudoku [][]int) {
	rows := sudoku
	columns := transpose_sudoku(sudoku)
	grids := grid_sudoku(sudoku)

	valid := true
	for _, row := range rows {
		valid_zone := check_zone((row))
		if valid_zone == false {
			valid = false
		}
	}

	for _, column := range columns {
		valid_zone := check_zone((column))
		if valid_zone == false {
			valid = false
		}
	}

	for _, grid := range grids {
		valid_zone := check_zone((grid))
		if valid_zone == false {
			valid = false
		}
	}

	if valid {
		fmt.Println("Valid!")
	} else {
		fmt.Println("Not valid :(")
	}
}

func check_zone(zone []int) bool {
	zoneSorted := append([]int{}, zone...)
	sort.Slice(zoneSorted, func(a, b int) bool {
		return zoneSorted[a] < zoneSorted[b]
	})
	return reflect.DeepEqual(zoneSorted, []int{1, 2, 3, 4, 5, 6, 7, 8, 9})
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
