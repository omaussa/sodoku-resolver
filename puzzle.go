package main

import "fmt"

type Sudoku struct {
	squares [][]*Square
	boxes   []*Box
}

func UpdateSudoku(sudoku [][]*Square, row int, column int) {
	number := sudoku[row][column].number
	for x := 0; x < SIZE_ROWS; x++ {
		if sudoku[x][column].possible[number-1] == 0 {
			sudoku[x][column].solvable--
		}
		sudoku[x][column].possible[number-1] = 1
	}

	for x := 0; x < SIZE_COLUMNS; x++ {
		if sudoku[row][x].possible[number-1] == 0 {
			sudoku[row][x].solvable--
		}
		sudoku[row][x].possible[number-1] = 1
	}
}

func (s *Sudoku) Print() {
	horizontalBars := "-------------------------------"
	fmt.Printf("%s\n", horizontalBars)
	for i := 0; i < SIZE_ROWS; i++ {
		fmt.Printf("|")
		for j := 0; j < SIZE_COLUMNS; j++ {
			fmt.Printf(" %d ", s.squares[i][j].number)
			if (j+1)%3 == 0 {
				fmt.Printf("|")
			}
		}
		fmt.Printf("\n")
		if (i+1)%3 == 0 {
			fmt.Printf("%s\n", horizontalBars)
		}
	}
}

func SetUpPuzzle(puzzle [][]int) *Sudoku {
	squares := make([][]*Square, SIZE_ROWS)
	boxes := CreateBoxes()
	currentBox := 0

	for i := 0; i < SIZE_ROWS; i++ {
		squares[i] = make([]*Square, SIZE_COLUMNS)
		for j := 0; j < SIZE_COLUMNS; j++ {
			squares[i][j] = &Square{
				number:   puzzle[i][j],
				row:      i,
				column:   j,
				solvable: 9,
				possible: [9]int{0},
			}
			boxes[currentBox].squares[boxes[currentBox].numbers] = squares[i][j]
			squares[i][j].box = boxes[currentBox]
			boxes[currentBox].numbers++
			if j == 2 || j == 5 {
				currentBox++
			}
		}
		currentBox -= 2
		if i == 2 {
			currentBox = 3
		}
		if i == 5 {
			currentBox = 6
		}
	}

	for i := 0; i < SIZE_ROWS; i++ {
		for j := 0; j < SIZE_ROWS; j++ {
			if squares[i][j].number != 0 {
				squares[i][j].solvable = 0
				UpdateSudoku(squares, i, j)
				UpdateBoxes(squares, i, j)
				UNSOLVED--
			}
		}
	}

	return &Sudoku{
		squares: squares,
		boxes:   boxes,
	}

}

func CreatePuzzle() [][]int {
	puzzle := [][]int{
		{0, 1, 9, 0, 0, 2, 0, 0, 0},
		{4, 7, 0, 6, 9, 0, 0, 0, 1},
		{0, 0, 0, 4, 0, 0, 0, 9, 0},

		{8, 9, 4, 5, 0, 7, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 2, 0, 1, 9, 5, 8},

		{0, 5, 0, 0, 0, 6, 0, 0, 0},
		{6, 0, 0, 0, 2, 8, 0, 7, 9},
		{0, 0, 0, 1, 0, 0, 8, 6, 0},
	}
	return puzzle
}

func CheckPuzzle(sudoku *Sudoku) int {
	squares := sudoku.squares
	for i := 0; i < SIZE_ROWS; i++ {
		for j := 0; j < SIZE_COLUMNS; j++ {
			if squares[i][j].solvable == 1 {
				squares[i][j].Solve()
				UpdateSudoku(squares, i, j)
				UpdateBoxes(squares, i, j)
				return 1
			}
		}
	}

	if BoxSingles(sudoku.squares, sudoku.boxes) == 1 {
		return 1
	}

	return CheckRows(sudoku.squares)
}
