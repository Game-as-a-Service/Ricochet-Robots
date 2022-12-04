package gamestatemanager

import (
	"ricochetrobots/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMoveUpward(t *testing.T) {
	// given
	bsm := NewRicochetRobotsManager(mockBoardManager())
	// when
	end := bsm.MoveRobot(models.Red, models.Upward)
	// then
	assert.Equal(t, models.NewPosition(0, 1), end)
}

func TestMoveRightward(t *testing.T) {
	// given
	bsm := NewRicochetRobotsManager(mockBoardManager())
	// when
	end := bsm.MoveRobot(models.Red, models.Downward)
	// then
	assert.Equal(t, models.NewPosition(3, 1), end)
}

func TestMoveDownward(t *testing.T) {
	// given
	bsm := NewRicochetRobotsManager(mockBoardManager())
	// when
	end := bsm.MoveRobot(models.Red, models.Leftward)
	// then
	assert.Equal(t, models.NewPosition(1, 0), end)
}

func TestMoveLeftward(t *testing.T) {
	// given
	bsm := NewRicochetRobotsManager(mockBoardManager())
	// when
	end := bsm.MoveRobot(models.Red, models.Rightward)
	// then
	assert.Equal(t, models.NewPosition(1, 3), end)
}

// func TestBlockedByWall(t *testing.T) {
// 	// given
// 	// when
// 	// then
// 	assert.True(t, false)
// }
// func TestBlockedByRobot(t *testing.T) {
// 	// given
// 	// when
// 	// then
// 	assert.True(t, false)
// }

func mockBoardManager() *models.BoardManager {
	return models.NewBoardManager(mockBoard())
}

func mockBoard() *models.Board {
	board := models.NewBoard(4)
	redRobotStartPoint := models.NewPosition(1, 1)
	board.SetRobot(models.Red, redRobotStartPoint) // TODO 這個有點越界了，外界不應該直接操作 board
	return board
}
