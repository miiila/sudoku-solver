package recursivesolver

import (
	"sudoku-solver/grid"
)

type sudoku = grid.Sudoku

type RecursiveSolver struct{}

func (rs RecursiveSolver) Solve(s sudoku) sudoku {
	numbers := [9]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	grid, _ := solve(s, numbers)

	return grid
}

func solve(s sudoku, numbers [9]int) (sudoku, bool) {
	nextCell, solved := grid.GetNextEmptyCell(s)
	if solved {
		return s, true
	}

	for _, v := range numbers {
		if grid.IsNumOK(s, v, nextCell) {
			s[nextCell.Row][nextCell.Col] = v
			g, solved := solve(s, numbers)
			if solved {
				return g, true
			}
		}
	}

	return s, false

}
