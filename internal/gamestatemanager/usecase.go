package gamestatemanager

import "ricochetrobots/models"

type GameStateManager interface {
	SetRobotOn(c string, p models.Position)
	MoveRobot(color, direction string) models.Position
}
