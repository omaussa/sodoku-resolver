package main

import "fmt"

func main() {
	puzzle := CreatePuzzle()
	var progress int

	sudoku := SetUpPuzzle(puzzle)

	sudoku.Print()

	for UNSOLVED > 0 {
		progress = CheckPuzzle(sudoku)
		if progress == 0 {
			fmt.Println("Failed to solve the puzzle!")
			break
		}
	}

	fmt.Printf("\n\n")

	sudoku.Print()

}
