package gamestatemanager

import "ricochetrobots/models"

type GameStateManager interface {
	MoveRobot(color, direction string) models.Position
}
