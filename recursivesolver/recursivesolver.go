package recursivesolver

import (
	"sudoku-solver/grid"
	"sync"
)

type sudoku = grid.Sudoku

type RecursiveSolver struct{}

var numbers = [9]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
func (rs RecursiveSolver) Solve(s sudoku) sudoku {
	grid, _ := solve(s)

	return grid
}

func (rs RecursiveSolver) SolveConcurrently(s sudoku, wg *sync.WaitGroup, res chan sudoku) {
	grid, _ := solve(s)

	res <- grid
	wg.Done()
}

func solve(s sudoku) (sudoku, bool) {
	nextCell, solved := grid.GetNextEmptyCell(s)
	if solved {
		return s, true
	}

	for _, v := range numbers {
		if grid.IsNumOK(s, v, nextCell) {
			s[nextCell.Row][nextCell.Col] = v
			g, solved := solve(s)
			if solved {
				return g, true
			}
		}
	}

	return s, false
}
