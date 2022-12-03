package gamestatemanager

import "ricochetrobots/models"

type GameStateManager interface {
	MoveRobot(c models.Color, d models.Direction) models.Position
}
