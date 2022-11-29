package gamestatemanager

import (
	"ricochetrobots/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMoveUpward(t *testing.T) {
	// given
	bsm := NewRicochetRobotsStateManager(NewBoardManager())
	bsm.SetRobotOn(models.RedRobot, models.NewPosition(2, 2))
	// when
	end := bsm.MoveRobot(models.RedRobot, models.Upward)
	// then
	assert.Equal(t, models.NewPosition(1, 2), end)
}
func TestMoveRightward(t *testing.T) {
	// given
	// when
	// then
	assert.True(t, false)
}
func TestMoveDownward(t *testing.T) {
	// given
	// when
	// then
	assert.True(t, false)
}
func TestMoveLeftward(t *testing.T) {
	// given
	// when
	// then
	assert.True(t, false)
}
func TestBlockedByWall(t *testing.T) {
	// given
	// when
	// then
	assert.True(t, false)
}
func TestBlockedByRobot(t *testing.T) {
	// given
	// when
	// then
	assert.True(t, false)
}
