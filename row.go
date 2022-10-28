package main

func CheckRows(sudoku [][]*Square) int {
	var place, sum [9]int
	for i := 0; i < SIZE_ROWS; i++ {
		sum = [9]int{0}
		place = [9]int{0}
		for j := 0; j < SIZE_ROWS; j++ {
			if sudoku[i][j].number != 0 {
				continue
			}
			for k := 0; k < 9; k++ {
				if sudoku[i][j].possible[k] == 0 {
					sum[k]++
					place[k] = j
				}
			}
		}
		for j := 0; j < SIZE_ROWS; j++ {
			if sum[j] == 1 {
				sudoku[i][place[j]].number = j + 1
				sudoku[i][place[j]].solvable = 0
				UNSOLVED--

				UpdateSudoku(sudoku, i, place[j])
				UpdateBoxes(sudoku, i, place[j])

				return 1
			}
		}
	}
	return 0
}
