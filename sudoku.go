package gokken

func findNextEmptyCell(grid [][]int) (int, int) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if grid[i][j] == 0 {
				return i, j
			}
		}
	}
	return -1, -1
}

func isPossible(grid [][]int, y, x, number int) bool {
	for i := 0; i < 9; i++ {
		if grid[y][i] == number || grid[i][x] == number {
			return false
		}
	}

	row := (y / 3) * 3
	col := (x / 3) * 3

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if grid[i+row][j+col] == number {
				return false
			}
		}
	}

	return true
}

func Solve(grid [][]int) bool {
	row, col := findNextEmptyCell(grid)

	if row == -1 && col == -1 {
		return true
	}

	for k := 1; k < 10; k++ {
		if isPossible(grid, row, col, k) {
			grid[row][col] = k
			if Solve(grid) {
				return true
			}
			grid[row][col] = 0
		}
	}

	return false
}
