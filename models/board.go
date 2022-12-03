package models

type Board struct {
	// 座標索引為(0,0) ~ (size-1, size-1)
	size int
	// 紀錄每一個 field 上的物件
	fields [][]Object
	// 紀錄所有機器人的位置
	robots map[Color]*Robot
}

func NewBoard(size int) *Board {
	fields := make([][]Object, size)
	for i := 0; i < size; i++ {
		fields[i] = make([]Object, size)
	}

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			fields[i][j] = NewGrid(0)
		}
	}
	// robot has fixed start position
	robots := make(map[Color]*Robot)
	robots[Red] = NewRobot(Position{})

	return &Board{size: size, fields: fields, robots: robots}
}

func (b Board) get(p Position) Object {
	return b.fields[p.row][p.col]
}

func (b Board) getRobot(c Color) *Robot {
	return b.robots[c]
}

func (b Board) isInBound(p Position) bool {
	return (p.row >= 0 && p.row < b.size) && (p.col >= 0 && p.col < b.size)
}

func (b *Board) SetRobot(c Color, new Position) {
	// record robot info
	b.robots[c].SetPosition(new)
	// update board
	old := b.robots[c].Position()
	b.set(new, b.robots[c])
	b.set(old, NewGrid(0)) // TODO 原本的地塊不一定是空地（這裡也引申出一個問題，你的機器人會覆蓋掉原本的地塊資訊）
}

func (b *Board) set(p Position, o Object) {
	b.fields[p.row][p.col] = o
}
