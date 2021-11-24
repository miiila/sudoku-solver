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

func main() {
	f, _ := os.Open("sudoku.txt")
	s := bufio.NewScanner(f)

	var p sudokusolver.ProjectEulerSudokuSolver
	p.Inputs = p.Parse(s, gridsInTxt, amount)
	p.Solver = recursivesolver.RecursiveSolver{}
	run(p.SolveIteratively, "RecursiveSolver iteratively - %d in %f seconds\n")
	run(p.SolveConcurrently, "RecursiveSolver concurrently - %d in %f seconds\n")

	p.Solver = routinessolver.RoutinesSolver{}
	run(p.SolveIteratively, "RoutinesSolver iteratively - %d in %f seconds\n")
	run(p.SolveConcurrently, "RoutinesSolver concurrently - %d in %f seconds\n")
}

func run(solver func() int, msg string) {
	start := time.Now()
	res := solver()
	fmt.Printf(msg, res, time.Since(start).Seconds())
}
