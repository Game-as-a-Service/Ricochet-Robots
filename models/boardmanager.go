package models

type BoardManager struct {
	board *Board
}

func NewBoardManager(b *Board) *BoardManager {
	bm := &BoardManager{
		board: b,
	}
	return bm
}

func (bm *BoardManager) Move(c Color, d Direction) Position {
	robot := bm.board.getRobot(c)
	cur := robot.Position()

	for bm.board.isInBound(cur.Next(d)) {
		next := cur.Next(d)
		// TODO 這個 IsBlock 也引用的太深了，要再思考一下這是誰的職責。
		if bm.board.get(next).IsBlock(d) {
			break
		}
		cur = next
	}

	robot.SetPosition(cur)
	return cur
}
