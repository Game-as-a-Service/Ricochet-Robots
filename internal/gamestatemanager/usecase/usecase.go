package gamestatemanager

import "ricochetrobots/models"

type RicochetRobotsStateManager struct {
	boardManager *BoardManager
}

func NewRicochetRobotsStateManager(bm *BoardManager) *RicochetRobotsStateManager {
	return &RicochetRobotsStateManager{bm}
}

func (sm *RicochetRobotsStateManager) SetRobotOn(c string, p models.Position) {
	sm.boardManager.SetRobotPosition(c, p)
}

func (sm *RicochetRobotsStateManager) MoveRobot(c, d string) models.Position {
	from := sm.boardManager.GetRobotPosition(models.RedRobot)
	return sm.boardManager.Move(from, models.Upward)
}

type BoardManager struct {
	board          *models.Board
	robotPositions map[string]models.Position
}

func NewBoardManager() *BoardManager {
	return &BoardManager{
		board: models.NewBoard(4),
		robotPositions: make(map[string]models.Position),
	}
}

func (bm *BoardManager) SetRobotPosition(c string, p models.Position) {
	bm.robotPositions[c] = p
	bm.board.Set(models.NewRobot(c), p)
}

func (bm *BoardManager) GetRobotPosition(c string) models.Position {
	return bm.robotPositions[c]
}

func (bm *BoardManager) Set(c string, p models.Position) {
	bm.board.Set(models.NewRobot(c), p)
}

func (bm *BoardManager) Move(from models.Position, to string) models.Position {

	current := from
	for current.Up() != models.NewPosition(0, 0) {
		if to == models.Upward {
			if bm.board.Get(current.Up()).BlockDown() {
				return current
			} else if bm.board.Get(current.Up()).BlockTop() {
				return current.Up()
			} else {
				current = current.Up()
			}
		}
	}
	return current
}
