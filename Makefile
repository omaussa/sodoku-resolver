all: sudoku run clean

sudoku:
	go build

run:
	./sudoku

clean:
	rm ./sudoku
