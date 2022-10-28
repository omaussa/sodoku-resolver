package main

type Square struct {
	number   int
	possible [SIZE_ROWS]int
	solvable int
	box      *Box
	row      int
	column   int
}

func (s *Square) Solve() {
	for i := 0; i < SIZE_ROWS; i++ {
		if s.possible[i] == 0 {
			s.number = i + 1
			s.solvable = 0
			UNSOLVED--
		}
	}
}
