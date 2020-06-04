package toyrobot

type State struct {
	MinX  int
	MaxX  int
	MinY  int
	MaxY  int
	Robot Robot
}

type Action struct {
	ActionType ActionType
	X          int
	Y          int
	Direction  Direction
}

type Robot struct {
	Position  Position
	Direction Direction
}

type Position struct {
	X int
	Y int
}

type Direction int
type ActionType string
type Message string

const (
	NORTH             Direction  = 0
	EAST              Direction  = 90
	SOUTH             Direction  = 180
	WEST              Direction  = 270
	PLACE             ActionType = "PLACE"
	MOVE              ActionType = "MOVE"
	LEFT              ActionType = "LEFT"
	RIGHT             ActionType = "RIGHT"
	REPORT            ActionType = "REPORT"
	InvalidActionType Message    = "Invalid action '%s', please check your input."
)
