package toyrobot

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestReducers(t *testing.T) {
	RegisterFailHandler(Fail)
}

var _ = Describe("Reducer", func() {
	var action Action
	var initialState State
	var finalState State
	var err error
	var position Position

	BeforeEach(func() {
		initialState = NewState()
	})

	//	Describe("reduces MAP action", func() {
	//		BeforeEach(func() {
	//			action = Action{ActionType: MAP}
	//			initialState.Robot.Position = Position{X: 0, Y: 1}
	//			initialState.Obstacles = append(initialState.Obstacles, Position{X: 4, Y: 3})
	//			finalState, _ = Reduce(action, initialState)
	//		})
	//
	//		It("generates a map", func() {
	//			Expect(finalState.Map).To(Equal(`
	//🏳🏳🏳🏳🏳
	//🏳🏳🏳🏳🏔
	//🏳🏳🏳🏳🏳
	//🤖🏳🏳🏳🏳
	//🏳🏳🏳🏳🏳
	//`))
	//		})
	//	})

	Describe("reduces PLACE_OBJECT action", func() {
		BeforeEach(func() {
			action = Action{ActionType: PLACE_OBJECT}
			finalState, _ = Reduce(action, initialState)
		})

		It("creates a new obstacle in the state", func() {
			Expect(len(finalState.Obstacles)).To(Equal(1))
		})

		Context("when the robot is on the border", func() {
			BeforeEach(func() {
				initialState.Robot.Direction = SOUTH
				finalState, _ = Reduce(action, initialState)
			})

			It("does NOT createa a new obstacle in the state", func() {
				Expect(len(finalState.Obstacles)).To(Equal(0))
			})
		})
	})

	Describe("reduces PLACE action", func() {
		BeforeEach(func() {
			action = Action{ActionType: PLACE, X: 3, Y: 2, Direction: EAST}
			finalState, _ = Reduce(action, initialState)
		})

		It("places the robot at the given position", func() {
			Expect(finalState.Robot.Position).To(Equal(Position{X: 3, Y: 2}))
		})

		It("rotates the robot facing the given direction", func() {
			Expect(finalState.Robot.Direction).To(Equal(EAST))
		})
	})

	Describe("reduces REPORT action", func() {
		BeforeEach(func() {
			action = Action{ActionType: REPORT}
			finalState, _ = Reduce(action, initialState)
		})

		It("generates a report string for the robot", func() {
			Expect(finalState.Report).To(Equal("The robot is at (0, 0) and it is facing North."))
		})

		Context("and when the robot is facing EAST", func() {
			BeforeEach(func() {
				initialState.Robot.Direction = EAST
				finalState, _ = Reduce(action, initialState)
			})

			It("generates a report string for the robot", func() {
				Expect(finalState.Report).To(Equal("The robot is at (0, 0) and it is facing East."))
			})
		})

		Context("and when the robot is facing SOUTH", func() {
			BeforeEach(func() {
				initialState.Robot.Direction = SOUTH
				finalState, _ = Reduce(action, initialState)
			})

			It("generates a report string for the robot", func() {
				Expect(finalState.Report).To(Equal("The robot is at (0, 0) and it is facing South."))
			})
		})

		Context("and when the robot is facing WEST", func() {
			BeforeEach(func() {
				initialState.Robot.Direction = WEST
				finalState, _ = Reduce(action, initialState)
			})

			It("generates a report string for the robot", func() {
				Expect(finalState.Report).To(Equal("The robot is at (0, 0) and it is facing West."))
			})
		})
	})

	Describe("reduces MOVE action", func() {
		BeforeEach(func() {
			action = Action{ActionType: MOVE}
		})

		Context("when there is ab obstacle in front of the robot", func() {
			BeforeEach(func() {
				initialState.Robot = Robot{Position: Position{X: 0, Y: 0}, Direction: NORTH}
				initialState.Obstacles = append(initialState.Obstacles, Position{X: 0, Y: 1})
				finalState, _ = Reduce(action, initialState)
			})

			It("does NOT move the robot", func() {
				Expect(finalState.Robot.Position).To(Equal(Position{X: 0, Y: 0}))
			})
		})

		Context("when the robot is facing SOUTH", func() {
			BeforeEach(func() {
				initialState.Robot.Position = Position{X: 3, Y: 2}
				initialState.Robot.Direction = SOUTH
				finalState, _ = Reduce(action, initialState)
			})

			It("moves the robot to the position in front of it", func() {
				Expect(finalState.Robot.Position).To(Equal(Position{X: 3, Y: 1}))
			})
		})

		Context("when the robot is facing SOUTH and it's on the border", func() {
			BeforeEach(func() {
				initialState.Robot.Direction = SOUTH
				finalState, _ = Reduce(action, initialState)
			})

			It("does NOT move the robot", func() {
				Expect(finalState.Robot.Position).To(Equal(Position{X: 0, Y: 0}))
			})
		})

		Context("when the robot is facing NORTH", func() {
			BeforeEach(func() {
				finalState, _ = Reduce(action, initialState)
			})

			It("moves the robot to the position in front of it", func() {
				Expect(finalState.Robot.Position).To(Equal(Position{X: 0, Y: 1}))
			})
		})

		Context("when the robot is facing NORTH and it's on the border", func() {
			BeforeEach(func() {
				initialState.Robot.Position = Position{X: 0, Y: 4}
				finalState, _ = Reduce(action, initialState)
			})

			It("does NOT move the robot", func() {
				Expect(finalState.Robot.Position).To(Equal(Position{X: 0, Y: 4}))
			})
		})

		Context("when the robot is facing EAST", func() {
			BeforeEach(func() {
				initialState.Robot.Direction = EAST
				initialState.Robot.Position = Position{X: 3, Y: 2}
				finalState, _ = Reduce(action, initialState)
			})

			It("moves the robot to the position in front of it", func() {
				Expect(finalState.Robot.Position).To(Equal(Position{X: 4, Y: 2}))
			})
		})

		Context("when the robot is facing EAST and it's on the border", func() {
			BeforeEach(func() {
				position = Position{X: 4, Y: 4}
				initialState = State{Robot: Robot{Position: position, Direction: EAST}, MaxX: 4, MaxY: 4}
				finalState, _ = Reduce(action, initialState)
			})

			It("does NOT move the robot", func() {
				Expect(finalState.Robot.Position).To(Equal(Position{X: 4, Y: 4}))
			})
		})

		Context("when the robot is facing WEST", func() {
			BeforeEach(func() {
				initialState.Robot.Direction = WEST
				initialState.Robot.Position = Position{X: 3, Y: 2}
				finalState, _ = Reduce(action, initialState)
			})

			It("moves the robot to the position in front of it", func() {
				Expect(finalState.Robot.Position).To(Equal(Position{X: 2, Y: 2}))
			})
		})

		Context("when the robot is facing WEST and it's on the border", func() {
			BeforeEach(func() {
				position = Position{X: 0, Y: 4}
				initialState = State{Robot: Robot{Position: position, Direction: WEST}, MaxX: 4, MaxY: 4}
				finalState, _ = Reduce(action, initialState)
			})

			It("does NOT move the robot", func() {
				Expect(finalState.Robot.Position).To(Equal(Position{X: 0, Y: 4}))
			})
		})

		Context("when the robot is facing an unknown direction", func() {
			BeforeEach(func() {
				initialState.Robot.Direction = 42
				initialState.Robot.Position = Position{X: 3, Y: 2}
				finalState, _ = Reduce(action, initialState)
			})

			It("does NOT move the robot", func() {
				Expect(finalState.Robot.Position).To(Equal(Position{X: 3, Y: 2}))
			})
		})
	})

	Describe("reduces LEFT action", func() {
		BeforeEach(func() {
			action = Action{ActionType: LEFT}
		})

		Context("when the robot is facing NORTH", func() {
			BeforeEach(func() {
				finalState, _ = Reduce(action, initialState)
			})

			It("rotates the robot to EAST", func() {
				Expect(finalState.Robot.Direction).To(Equal(WEST))
			})
		})

		Context("when the robot is facing SOUTH", func() {
			BeforeEach(func() {
				initialState.Robot.Direction = SOUTH
				finalState, _ = Reduce(action, initialState)
			})

			It("rotates the robot to EAST", func() {
				Expect(finalState.Robot.Direction).To(Equal(EAST))
			})
		})

		Context("when the robot is facing WEST", func() {
			BeforeEach(func() {
				initialState.Robot.Direction = WEST
				finalState, _ = Reduce(action, initialState)
			})

			It("rotates the robot to SOUTH", func() {
				Expect(finalState.Robot.Direction).To(Equal(SOUTH))
			})
		})

		Context("when the robot is facing EAST", func() {
			BeforeEach(func() {
				initialState.Robot.Direction = EAST
				finalState, _ = Reduce(action, initialState)
			})

			It("rotates the robot to SOUTH", func() {
				Expect(finalState.Robot.Direction).To(Equal(NORTH))
			})
		})
	})

	Describe("reduces RIGHT action", func() {
		BeforeEach(func() {
			action = Action{ActionType: RIGHT}
		})

		Context("when the robot is facing NORTH", func() {
			BeforeEach(func() {
				finalState, _ = Reduce(action, initialState)
			})

			It("rotates the robot to EAST", func() {
				Expect(finalState.Robot.Direction).To(Equal(EAST))
			})
		})

		Context("when the robot is facing SOUTH", func() {
			BeforeEach(func() {
				initialState.Robot.Direction = SOUTH
				finalState, _ = Reduce(action, initialState)
			})

			It("rotates the robot to EAST", func() {
				Expect(finalState.Robot.Direction).To(Equal(WEST))
			})
		})

		Context("when the robot is facing WEST", func() {
			BeforeEach(func() {
				initialState.Robot.Direction = WEST
				finalState, _ = Reduce(action, initialState)
			})

			It("rotates the robot to SOUTH", func() {
				Expect(finalState.Robot.Direction).To(Equal(NORTH))
			})
		})

		Context("when the robot is facing EAST", func() {
			BeforeEach(func() {
				initialState.Robot.Direction = EAST
				finalState, _ = Reduce(action, initialState)
			})

			It("rotates the robot to SOUTH", func() {
				Expect(finalState.Robot.Direction).To(Equal(SOUTH))
			})
		})
	})

	Context("reduces unknown actions", func() {
		Context("when the action is not recognised", func() {
			BeforeEach(func() {
				action = Action{ActionType: "SPAM"}
				_, err = Reduce(action, initialState)
			})

			It("returns an error", func() {
				Expect(err.Error()).To(Equal("Invalid action 'SPAM', please check your input."))
			})
		})
	})
})
