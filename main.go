package main

import "fmt"

func main() {
	puzzle := CreatePuzzle()

	sudoku := SetUpPuzzle(puzzle)

	sudoku.Print()

	CheckPuzzle(sudoku)

	fmt.Printf("\n\n")

	sudoku.Print()

}
