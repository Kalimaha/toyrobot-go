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
		finalState.Robot.Direction = initialState.Robot.Direction - 90
		if finalState.Robot.Direction < 0 {
			finalState.Robot.Direction += 360
		}
		return finalState, nil
	case RIGHT:
		finalState.Robot.Direction = initialState.Robot.Direction + 90
		if finalState.Robot.Direction == 360 {
			finalState.Robot.Direction = 0
		}
		return finalState, nil
	default:
		return finalState, errors.New(fmt.Sprintf(string(INVALID_ACTION_TYPE), action.ActionType))
	}
}
