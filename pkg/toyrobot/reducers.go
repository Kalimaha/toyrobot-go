package toyrobot

import (
	"errors"
	"fmt"
	"github.com/jinzhu/copier"
)

func Reduce(action Action, initialState State) (finalState State, err error) {
	_ = copier.Copy(&finalState, &initialState)

	switch action.ActionType {
	case MOVE:
		return move(finalState)
	case LEFT:
		return rotateLeft(finalState)
	case RIGHT:
		return rotateRight(finalState)
	case REPORT:
		return report(finalState)
	default:
		return finalState, errors.New(fmt.Sprintf(string(InvalidActionType), action.ActionType))
	}
}

func report(state State) (State, error) {
	state.Report = fmt.Sprintf(
		"The robot is at (%d, %d) and it is facing %s.",
		state.Robot.Position.X,
		state.Robot.Position.Y,
		int2direction(state.Robot.Direction))

	return state, nil
}

func int2direction(d Direction) string {
	switch d {
	case 90:
		return "East"
	case 180:
		return "South"
	case 270:
		return "West"
	default:
		return "North"
	}
}

func move(state State) (State, error) {
	newPosition := GetNextPosition(state.Robot)
	if isValidPosition(state, newPosition) {
		state.Robot.Position = newPosition
	}
	return state, nil
}

func GetNextPosition(robot Robot) Position {
	switch robot.Direction {
	case NORTH:
		return Position{X: robot.Position.X, Y: robot.Position.Y + 1}
	case SOUTH:
		return Position{X: robot.Position.X, Y: robot.Position.Y - 1}
	case EAST:
		return Position{X: robot.Position.X + 1, Y: robot.Position.Y}
	case WEST:
		return Position{X: robot.Position.X - 1, Y: robot.Position.Y}
	default:
		return robot.Position
	}
}

func isValidPosition(state State, position Position) bool {
	if position.X < state.MinX || position.X > state.MaxX || position.Y < state.MinY || position.Y > state.MaxY {
		return false
	}
	return true
}

func rotateLeft(state State) (State, error) {
	state.Robot.Direction = state.Robot.Direction - 90
	if state.Robot.Direction < 0 {
		state.Robot.Direction += 360
	}
	return state, nil
}

func rotateRight(state State) (State, error) {
	state.Robot.Direction = state.Robot.Direction + 90
	if state.Robot.Direction == 360 {
		state.Robot.Direction = 0
	}
	return state, nil
}
