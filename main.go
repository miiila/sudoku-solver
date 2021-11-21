package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sudoku-solver/grid"
	"sudoku-solver/recursivesolver"
)

type sudoku = grid.Sudoku

const gridsInTxt = 50
const amount = 1000
const ratio = amount / gridsInTxt

func main() {
	f, _ := os.Open("sudoku.txt")
	s := bufio.NewScanner(f)

	var p ProjectEulerSudokuSolver
	p.inputs = p.parse(s)
	p.solver = recursivesolver.RecursiveSolver{}

	fmt.Println(p.solve())
}

type ProjectEulerInputParser interface {
	Parse(*bufio.Scanner) [amount]sudoku
}

type ProjectEulerSudokuSolver struct {
	inputs [amount]sudoku
	solver SudokuSolver
}

type SudokuSolver interface {
	Solve(s sudoku) sudoku
}

func (p ProjectEulerSudokuSolver) parse(scanner *bufio.Scanner) [amount]sudoku {
	i := -1
	var j int
	var inputs [amount]sudoku
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

func (p ProjectEulerSudokuSolver) solve() int {
	var res int
	for _, s := range p.inputs {
		solved := p.solver.Solve(s)
		res += solved[0][0]*100 + solved[0][1]*10 + solved[0][2]
	}

	return res
}
