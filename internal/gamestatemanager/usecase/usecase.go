package gamestatemanager

import "ricochetrobots/models"

type RicochetRobotsStateManager struct {
}

func (sm *RicochetRobotsStateManager) MoveRobot(color, direction string) models.Position {
	return models.Position{}
}

type BoardManager struct {
	board Board
}

func (bm *BoardManager) GetRobotPosition(color string) models.Position {
	return models.Position{}
}
func (bm *BoardManager) Move(from models.Position, direction string) models.Position {
	return models.Position{}
}

type Board struct {
}
