package models

type Object interface {
	// 每一個塊都只負責判斷對應方向上的那堵牆，比如有一個來自上方的物件，我只判斷我有沒有上面那堵牆
	// 而不判斷我有沒有下面那堵牆，因為那是下面那一個塊地的職責。
	// 所以你在生成牆壁時，相鄰的兩個地塊都要生成那堵牆。
	IsBlock(Direction) bool
}

type Robot struct {
	position Position
}

func NewRobot(p Position) *Robot {
	return &Robot{p}
}

func (r Robot) IsBlock(d Direction) bool {
	return true
}

func (r *Robot) SetPosition(p Position) {
	r.position = p
}

func (r *Robot) Position() Position {
	return r.position
}

func (r *Robot) Row() int {
	return r.position.row
}

func (r *Robot) Col() int {
	return r.position.col
}


type Grid struct {
	walls int
}

func NewGrid(walls int) Grid {
	return Grid{walls}
}

func (g Grid) IsBlock(d Direction) bool {
	return g.walls&int(d) != 0
}
