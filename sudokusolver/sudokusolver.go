package sudokusolver

import (
	"bufio"
	"strings"
	"sudoku-solver/grid"
	"sync"
)

type sudoku = grid.Sudoku

type ProjectEulerInputParser interface {
	Parse(*bufio.Scanner) []sudoku
}

type ProjectEulerSudokuSolver struct {
	Inputs []sudoku
	Solver SudokuSolver
}

type SudokuSolver interface {
	Solve(s sudoku) sudoku
	SolveConcurrently(s sudoku, wg *sync.WaitGroup, res chan sudoku)
}

func (p ProjectEulerSudokuSolver) Parse(scanner *bufio.Scanner, ratio int, gridsInTxt int, amount int) []sudoku {
	i := -1
	var j int
	var inputs  = make([]sudoku, amount)
	var s sudoku
	for scanner.Scan() {
		t := scanner.Text()
		if strings.Contains(t, "Grid") {
			if i > -1 {
				for k := 0; k < ratio; k++ {
					inputs[i+(gridsInTxt*k)] = s
				}
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
	for k := 0; k < amount/gridsInTxt; k++ {
		inputs[i+(gridsInTxt*k)] = s
	}

	return inputs
}

func (p ProjectEulerSudokuSolver) SolveIteratively() int {
	var res int
	for _, s := range p.Inputs {
		solved := p.Solver.Solve(s)
		res += solved[0][0]*100 + solved[0][1]*10 + solved[0][2]
	}

	return res
}

func (p ProjectEulerSudokuSolver) SolveConcurrently() int {
	var res int
	var wg sync.WaitGroup
	resChan := make(chan sudoku, 1000)
	for _, s := range p.Inputs {
		wg.Add(1)
		go p.Solver.SolveConcurrently(s, &wg, resChan)
	}

	wg.Wait()
	close(resChan)
	for solved := range resChan {
		res += solved[0][0]*100 + solved[0][1]*10 + solved[0][2]
	}
	return res
}
