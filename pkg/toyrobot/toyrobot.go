package toyrobot

import "fmt"

func Play(filepath string, logLevel LogLevel) (state State, err error) {
	actions, err := ParseFile(filepath)
	if err != nil {
		fmt.Println(err.Error())
		return state, err
	} else {
		state := NewState()
		for _, action := range actions {
			state, err = Reduce(action, state)
			if err != nil {
				fmt.Println(err.Error())
				return state, err
			} else if action.ActionType == REPORT && logLevel == INFO {
				printReport(state)
			}
		}
		return state, err
	}
}

func printReport(state State) {
	fmt.Println()
	fmt.Println("Toy Robot Report")
	fmt.Println("----------------")
	fmt.Println(state.Report)
	fmt.Println()
}
