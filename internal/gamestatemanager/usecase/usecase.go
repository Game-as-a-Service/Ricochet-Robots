package gamestatemanager

import "ricochetrobots/models"

type RicochetRobotsManager struct {
	boardManager *models.BoardManager
}

func NewRicochetRobotsManager(bm *models.BoardManager) *RicochetRobotsManager {
	return &RicochetRobotsManager{
		boardManager: bm,
	}
}

func (rm *RicochetRobotsManager) MoveRobot(c models.Color, d models.Direction) models.Position {
	return rm.boardManager.Move(c, d)
}
