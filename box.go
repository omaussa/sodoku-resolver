package main

type Box struct {
	squares  []*Square
	numbers  int
	possible [9]int
	solvable int
	next     *Box
}

func CreateBoxes() []*Box {
	boxes := make([]*Box, 9)

	for i := 0; i < 9; i++ {
		boxes[i] = &Box{
			squares:  make([]*Square, 9),
			numbers:  0,
			solvable: 9,
			possible: [9]int{0},
		}
	}

	return boxes
}

func UpdateBoxes(sudoku [][]*Square, row int, column int) {
	number := sudoku[row][column].number
	box := sudoku[row][column].box
	for i := 0; i < 9; i++ {
		if box.squares[i].possible[number-1] == 0 {
			box.squares[i].solvable--
			box.squares[i].possible[number-1] = 1
		}
	}
}

func BoxSingles(sudoku [][]*Square, boxes []*Box) {
	temp := 0
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			counter := 0
			for x := 0; x < 9; x++ {
				if boxes[i].squares[x].number != 0 {
					continue
				}
				if boxes[i].squares[x].possible[j] == 0 {
					counter++
					temp = x
				}
				if counter == 2 {
					break
				}
			}
			if counter == 1 {
				boxes[i].squares[temp].number = j + 1
				UNSOLVED--
				boxes[i].squares[temp].solvable = 0

				updateSudoku(sudoku, boxes[i].squares[temp].row, boxes[i].squares[temp].column)
			}
		}
	}
}
