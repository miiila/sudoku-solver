package main

import (
	"bufio"
	"fmt"
	"os"
	"sudoku-solver/grid"
	"sudoku-solver/recursivesolver"
	"sudoku-solver/routinessolver"
	"sudoku-solver/sudokusolver"
	"time"
)

type sudoku = grid.Sudoku

const gridsInTxt = 50
const amount = 1000
const ratio = amount / gridsInTxt

func main() {
	f, _ := os.Open("sudoku.txt")
	s := bufio.NewScanner(f)

	var p sudokusolver.ProjectEulerSudokuSolver
    p.Inputs = p.Parse(s, ratio, gridsInTxt, amount)
	p.Solver = recursivesolver.RecursiveSolver{}
	startTime := time.Now()
	res := p.SolveIteratively()
	endTime := time.Now()
	diff := endTime.Sub(startTime)
	fmt.Printf("RecursiveSolver iteratively - %d in %f seconds\n", res, diff.Seconds())
	startTime = time.Now()
	res = p.SolveConcurrently()
	endTime = time.Now()
	diff = endTime.Sub(startTime)
	fmt.Printf("RecursiveSolver concurrently - %d in %f seconds\n", res, diff.Seconds())

	p.Solver = routinessolver.RoutinesSolver{}
	startTime = time.Now()
	res = p.SolveIteratively()
	endTime = time.Now()
	diff = endTime.Sub(startTime)

	fmt.Printf("RoutinesSolver iteratively - %d in %f seconds\n", res, diff.Seconds())
	startTime = time.Now()
	res = p.SolveConcurrently()
	endTime = time.Now()
	diff = endTime.Sub(startTime)

	fmt.Printf("RoutinesSolver concurrently - %d in %f seconds\n", res, diff.Seconds())
}
