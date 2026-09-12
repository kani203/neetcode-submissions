var keyCnt = make(map[int]int)

func isValidSudoku(board [][]byte) bool {
	// valid per row
	for i := 0; i < 9; i++ {
		isValidRow := isValid([]int{i, 0}, []int{i, 8}, board)
		if isValidRow == false {
			return false
		}
	}
	// valid per col
	for i := 0; i < 9; i++ {
		isValidCol := isValid([]int{0, i}, []int{8, i}, board)
		if isValidCol == false {
			return false
		}
	}
	// valid per grid
	for i := 0; i < 9; i += 3 {
		for j := 0; j < 9; j += 3 {
			isValidGrid := isValid([]int{i, j}, []int{i + 2, j + 2}, board)
			if isValidGrid == false {
				return false
			}
		}
	}
	return true
}

// valid from start(i,j) to end(m,n)
func isValid(start []int, end []int, board [][]byte) bool {
	resetCounter()
	for i := start[0]; i <= end[0]; i++ {
		for j := start[1]; j <= end[1]; j++ {
			if board[i][j] == '.' {
				continue
			}

			keyCnt[int(board[i][j])]++
			if keyCnt[int(board[i][j])] > 1 {
				fmt.Printf("start is %v, end is %v", start, end)
				fmt.Printf("i = %d, j = %d", i, j)
				return false
			}
		}
	}
	return true
}

func resetCounter() {
	for i := 1; i < 10; i++ {
		keyCnt[int(byte('0')+byte(i))] = 0
	}
}
