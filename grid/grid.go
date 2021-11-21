package grid

type Sudoku = [9][9]int

type Coord struct {
	Row int
	Col int
}

func GetNextEmptyCell(grid Sudoku) (Coord, bool) {
	for i, v := range grid {
		for j, w := range v {
			if w == 0 {
				return Coord{i, j}, false
			}
		}
	}
	return Coord{}, true
}

func IsNumOK(grid Sudoku, numberToFill int, CoordToFill Coord) bool {
	return isRowOK(grid, numberToFill, CoordToFill) && isColOK(grid, numberToFill, CoordToFill) && isSquareOK(grid, numberToFill, CoordToFill)
}

func isRowOK(grid Sudoku, numberToFill int, CoordToFill Coord) bool {
	for _, v := range grid[CoordToFill.Row] {
		if v == numberToFill {
			return false
		}
	}
	return true
}

func isColOK(grid Sudoku, numberToFill int, CoordToFill Coord) bool {
	for v := 0; v < 9; v++ {
		if grid[v][CoordToFill.Col] == numberToFill {
			return false
		}
	}
	return true
}

func isSquareOK(grid Sudoku, numberToFill int, CoordToFill Coord) bool {
	var squareStart Coord
	squareStart.Row = 3 * (CoordToFill.Row / 3)
	squareStart.Col = 3 * (CoordToFill.Col / 3)
	for i := squareStart.Row; i < squareStart.Row+3; i++ {
		for j := squareStart.Col; j < squareStart.Col+3; j++ {
			if grid[i][j] == numberToFill {
				return false
			}
		}
	}
	return true
}
