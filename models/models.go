package models

const (
	Upward    = "up"
	Downward  = "downward"
	Leftward  = "left"
	Rightward = "right"
)

const (
	RedRobot = "red"
)

type Board struct {
	size   int
	fields [][]Obstacle
}

func NewBoard(size int) *Board {
	fields := make([][]Obstacle, size)
	for i := 0; i < size; i++ {
		fields[i] = make([]Obstacle, size)
	}

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if i == 0 {
				fields[i][j] = NewWall(true, true, false, false)
			} else {
				fields[i][j] = NewWall(false, false, false, false)
			}
		}
	}

	return &Board{size: size, fields: fields}
}

func (b *Board) Get(p Position) Obstacle {
	return b.fields[p.x][p.y]
}

func (b *Board) Set(o Obstacle, p Position) {
	b.fields[p.x][p.y] = o
}

type Position struct {
	x, y int
}

func NewPosition(x, y int) Position {
	return Position{x, y}
}

func (p Position) Up() Position {
	if p.x == 0 {
		return Position{0, 0}
	}
	return Position{p.x-1, p.y}
}

type Obstacle interface {
	BlockTop() bool
	BlockDown() bool
}

type Wall struct {
	blockTop   bool
	blockDown  bool
	blockLeft  bool
	blockRight bool
}

func NewWall(blockTop, blockDown, blockLeft, blockRight bool) Wall {
	return Wall{blockTop, blockDown, blockLeft, blockRight}
}

func (w Wall) BlockTop() bool {
	return w.blockTop
}

func (w Wall) BlockDown() bool {
	return w.blockDown
}

type Robot struct {
	color string
}

func NewRobot(color string) Robot {
	return Robot{color}
}

func (r Robot) BlockTop() bool {
	return true
}

func (r Robot) BlockDown() bool {
	return true
}
