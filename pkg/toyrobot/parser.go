package toyrobot

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ParseFile(path string) (actions []Action, err error) {
	file, openErr := os.Open(path)
	if openErr != nil {
		return actions, errors.New("file does not exist")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		action, lineError := ParseLine(scanner.Text())
		if lineError == nil {
			actions = append(actions, action)
		}
	}

	return cleanActions(actions), err
}

func cleanActions(actions []Action) []Action {
	var cleanActions []Action
	placeActionFound := false

	for _, action := range actions {
		if placeActionFound == false && action.ActionType == PLACE {
			placeActionFound = true
		}

		if placeActionFound {
			cleanActions = append(cleanActions, action)
		}
	}

	return cleanActions
}

func ParseLine(s string) (action Action, err error) {
	if s == string(MOVE) {
		return Action{ActionType: MOVE}, err
	} else if s == string(LEFT) {
		return Action{ActionType: LEFT}, err
	} else if s == string(RIGHT) {
		return Action{ActionType: RIGHT}, err
	} else if s == string(REPORT) {
		return Action{ActionType: REPORT}, err
	} else if strings.HasPrefix(s, string(PLACE_OBJECT)) {
		return Action{ActionType: PLACE_OBJECT}, err
	} else if strings.HasPrefix(s, string(PLACE)) {
		return parsePlace(s), err
	} else {
		return action, errors.New(fmt.Sprintf("invalid action: %s", s))
	}
}

func parsePlace(s string) Action {
	parts := strings.Split(s, " ")
	coordinates := strings.Split(parts[len(parts)-1], ",")
	x, _ := strconv.Atoi(coordinates[0])
	y, _ := strconv.Atoi(coordinates[1])
	d, _ := str2dir(coordinates[2])

	return Action{ActionType: PLACE, X: x, Y: y, Direction: d}
}

func str2dir(s string) (direction Direction, err error) {
	switch s {
	case "NORTH":
		return NORTH, err
	case "EAST":
		return EAST, err
	case "WEST":
		return WEST, err
	case "SOUTH":
		return SOUTH, err
	default:
		return direction, errors.New(fmt.Sprintf("Unknown direction: %s", s))
	}
}
