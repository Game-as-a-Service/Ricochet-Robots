package models

type Direction int
const (
	Upward    Direction = 1 << iota
	Downward            = 1 << iota
	Leftward            = 1 << iota
	Rightward           = 1 << iota
)

type Color int
const (
	Red Color = iota
	Green
	Blue
	Yellow
)
