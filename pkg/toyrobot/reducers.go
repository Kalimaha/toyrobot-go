package toyrobot

import (
	"errors"
	"fmt"
	"github.com/jinzhu/copier"
)

func Reduce(action Action, initialState State) (finalState State, err error) {
	_ = copier.Copy(&finalState, &initialState)

	switch action.ActionType {
	case LEFT:
		return rotateLeft(finalState)
	case RIGHT:
		return rotateRight(finalState)
	default:
		return finalState, errors.New(fmt.Sprintf(string(InvalidActionType), action.ActionType))
	}
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
