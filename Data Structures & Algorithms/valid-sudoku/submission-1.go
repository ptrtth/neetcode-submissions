func uniqueNums(values []byte) bool {
	seen := make(map[byte]bool, 9)

	for _, value := range values {
		if value == '.' {
			continue
		}

		if seen[value] {
			return false
		}

		seen[value] = true
	}

	return true
}

func isValidSudoku(board [][]byte) bool {
	// Rows
	for _, row := range board {
		if !uniqueNums(row) {
			return false
		}
	}

	// Columns
	for col := 0; col < 9; col++ {
		values := make([]byte, 0, 9)

		for row := 0; row < 9; row++ {
			values = append(values, board[row][col])
		}

		if !uniqueNums(values) {
			return false
		}
	}

	// 3×3 boxes
	for boxRow := 0; boxRow < 9; boxRow += 3 {
		for boxCol := 0; boxCol < 9; boxCol += 3 {
			values := make([]byte, 0, 9)

			for row := boxRow; row < boxRow+3; row++ {
				for col := boxCol; col < boxCol+3; col++ {
					values = append(values, board[row][col])
				}
			}

			if !uniqueNums(values) {
				return false
			}
		}
	}

	return true
}