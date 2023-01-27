package main

func isSafe(m *[][]int8, row, col int) bool {
	for i := 0; i < len(*m); i++ {
		if (*m)[i][col] == 1 {
			return false
		}
	}

	var i, j int = row, col
	for i >= 0 && j >= 0 {
		if (*m)[i][j] == 1 {
			return false
		}
		i--
		j--
	}

	i, j = row, col
	for i >= 0 && j < len(*m) {
		if (*m)[i][j] == 1 {
			return false
		}
		i--
		j++
	}

	return true
}

func nQueen(m *[][]int8, row int) bool {
	if row >= len(*m) {
		return true
	}

	for i := 0; i < len(*m); i++ {
		if isSafe(m, row, i) {
			(*m)[row][i] = 1
			if nQueen(m, row+1) {
				return true
			}
			(*m)[row][i] = 0
		}
	}
	return false
}

func main() {
	example := [][]int8{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}

	nQueen(&example, 0)
	show(&example)
}
