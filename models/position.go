package models

var (
	rowStep = map[Direction]int{Upward:-1, Downward:1}
	colStep = map[Direction]int{Leftward:-1, Rightward:1}
)

type Position struct {
	row, col int
}

func NewPosition(row, col int) Position {
	return Position{row, col}
}

func (p Position) Next(d Direction) Position {
	return NewPosition(p.row+rowStep[d], p.col+colStep[d])
}
