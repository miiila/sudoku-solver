package routinessolver

import (
	"sudoku-solver/grid"
	"sync"
)

type sudoku = grid.Sudoku

type RoutinesSolver struct{}

var numbers = [9]int{1, 2, 3, 4, 5, 6, 7, 8, 9}

func (rs RoutinesSolver) Solve(s sudoku) sudoku {
	done := make(chan sudoku)
	go solve(s, done)

	grid := <-done

	return grid
}

func (rs RoutinesSolver) SolveConcurrently(s sudoku, wg *sync.WaitGroup, res chan sudoku) {
	done := make(chan sudoku)
	go solve(s, done)

	grid := <-done

	res <- grid
	wg.Done()
}

func solve(s sudoku, done chan sudoku) {
	nextCell, solved := grid.GetNextEmptyCell(s)
	if solved {
		done <- s
	}

	for _, v := range numbers {
		if grid.IsNumOK(s, v, nextCell) {
			s[nextCell.Row][nextCell.Col] = v
			go solve(s, done)
		}
	}
}

//func solver(grids chan sudoku, done chan sudoku) {
//for grids := range jobs {
//output := Result{job, digits(job.randomno)}
//results <- output
//}
//}
//}
