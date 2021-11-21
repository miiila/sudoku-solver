package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type sudoku = [9][9]int

type coord struct {
	row int
	col int
}

func main() {
	f, _ := os.Open("sudoku.txt")
	s := bufio.NewScanner(f)
	inputs := parse(s)

	var res int
	for _, s := range inputs {
		solved := solver(s)
		res += solved[0][0]*100 + solved[0][1]*10 + solved[0][2]
	}

	fmt.Println(res)
}

func solver(s sudoku) sudoku {
	numbers := [9]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	grid, _ := solveRec(s, numbers)

	return grid
}

func solveRec(grid sudoku, numbers [9]int) (sudoku, bool) {
	nextCell, solved := getNextEmptyCell(grid)
	if solved {
		return grid, true
	}

	for _, v := range numbers {
		if isNumOK(grid, v, nextCell) {
			grid[nextCell.row][nextCell.col] = v
			g, solved := solveRec(grid, numbers)
			if solved {
				return g, true
			}
		}
	}

	return grid, false

}

func getNextEmptyCell(grid sudoku) (coord, bool) {
	for i, v := range grid {
		for j, w := range v {
			if w == 0 {
				return coord{i, j}, false
			}
		}
	}
	return coord{}, true
}

func isNumOK(grid sudoku, numberToFill int, coordToFill coord) bool {
	return isRowOK(grid, numberToFill, coordToFill) && isColOK(grid, numberToFill, coordToFill) && isSquareOK(grid, numberToFill, coordToFill)
}

func isRowOK(grid sudoku, numberToFill int, coordToFill coord) bool {
	for _, v := range grid[coordToFill.row] {
		if v == numberToFill {
			return false
		}
	}
	return true
}

func isColOK(grid sudoku, numberToFill int, coordToFill coord) bool {
	for v := 0; v < 9; v++ {
		if grid[v][coordToFill.col] == numberToFill {
			return false
		}
	}
	return true
}

func isSquareOK(grid sudoku, numberToFill int, coordToFill coord) bool {
	var squareStart coord
	squareStart.row = 3 * (coordToFill.row / 3)
	squareStart.col = 3 * (coordToFill.col / 3)
	for i := squareStart.row; i < squareStart.row+3; i++ {
		for j := squareStart.col; j < squareStart.col+3; j++ {
			if grid[i][j] == numberToFill {
				return false
			}
		}
	}
	return true
}

func parse(scanner *bufio.Scanner) [50]sudoku {
	var inputs [50]sudoku
	i := -1
	var j int
	var s sudoku
	for scanner.Scan() {
		t := scanner.Text()
		if strings.Contains(t, "Grid") {
			if i > -1 {
				inputs[i] = s
			}
			i++
			j = 0
			continue
		}
		for k, v := range t {
			s[j][k] = int(v - '0')
		}
		j++
	}
	inputs[i] = s
	return inputs
}
