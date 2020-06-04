package toyrobot

import (
	"github.com/Kalimaha/ginkgo/reporter"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestReducers(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecsWithCustomReporters(t, "Toy Robot Reducer", []Reporter{reporter.New()})
}

var _ = Describe("Reducer", func() {
	var action Action
	var initialState State
	var finalState State
	var err error
	var position Position

	BeforeEach(func() {
		position = Position{X: 3, Y: 2}
		initialState = State{Robot: Robot{Position: position, Direction: SOUTH}, MaxX: 4, MaxY: 4}
	})

	Describe("reduces MOVE action", func() {
		BeforeEach(func() {
			action = Action{ActionType: MOVE}
		})

		Context("when the robot is facing SOUTH", func() {
			BeforeEach(func() {
				initialState = State{Robot: Robot{Position: position, Direction: SOUTH}, MaxX: 4, MaxY: 4}
				finalState, _ = Reduce(action, initialState)
			})

			It("moves the robot to the position in front of it", func() {
				Expect(finalState.Robot.Position).To(Equal(Position{X: 3, Y: 1}))
			})
		})

		Context("when the robot is facing SOUTH and it's on the border", func() {
			BeforeEach(func() {
				position = Position{X: 0, Y: 0}
				initialState = State{Robot: Robot{Position: position, Direction: SOUTH}, MaxX: 4, MaxY: 4}
				finalState, _ = Reduce(action, initialState)
			})

			It("does NOT move the robot", func() {
				Expect(finalState.Robot.Position).To(Equal(Position{X: 0, Y: 0}))
			})
		})

		Context("when the robot is facing NORTH", func() {
			BeforeEach(func() {
				initialState = State{Robot: Robot{Position: position, Direction: NORTH}, MaxX: 4, MaxY: 4}
				finalState, _ = Reduce(action, initialState)
			})

			It("moves the robot to the position in front of it", func() {
				Expect(finalState.Robot.Position).To(Equal(Position{X: 3, Y: 3}))
			})
		})

		Context("when the robot is facing NORTH and it's on the border", func() {
			BeforeEach(func() {
				position = Position{X: 0, Y: 4}
				initialState = State{Robot: Robot{Position: position, Direction: NORTH}, MaxX: 4, MaxY: 4}
				finalState, _ = Reduce(action, initialState)
			})

			It("does NOT move the robot", func() {
				Expect(finalState.Robot.Position).To(Equal(Position{X: 0, Y: 4}))
			})
		})

		Context("when the robot is facing EAST", func() {
			BeforeEach(func() {
				initialState = State{Robot: Robot{Position: position, Direction: EAST}, MaxX: 4, MaxY: 4}
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
				initialState = State{Robot: Robot{Position: position, Direction: WEST}, MaxX: 4, MaxY: 4}
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
	})

	Describe("reduces LEFT action", func() {
		BeforeEach(func() {
			action = Action{ActionType: LEFT}
		})

		Context("when the robot is facing NORTH", func() {
			BeforeEach(func() {
				initialState = State{Robot: Robot{Position: position, Direction: NORTH}, MaxX: 4, MaxY: 4}
				finalState, _ = Reduce(action, initialState)
			})

			It("rotates the robot to EAST", func() {
				Expect(finalState.Robot.Direction).To(Equal(WEST))
			})
		})

		Context("when the robot is facing SOUTH", func() {
			BeforeEach(func() {
				initialState = State{Robot: Robot{Position: position, Direction: SOUTH}, MaxX: 4, MaxY: 4}
				finalState, _ = Reduce(action, initialState)
			})

			It("rotates the robot to EAST", func() {
				Expect(finalState.Robot.Direction).To(Equal(EAST))
			})
		})

		Context("when the robot is facing WEST", func() {
			BeforeEach(func() {
				initialState = State{Robot: Robot{Position: position, Direction: WEST}, MaxX: 4, MaxY: 4}
				finalState, _ = Reduce(action, initialState)
			})

			It("rotates the robot to SOUTH", func() {
				Expect(finalState.Robot.Direction).To(Equal(SOUTH))
			})
		})

		Context("when the robot is facing EAST", func() {
			BeforeEach(func() {
				initialState = State{Robot: Robot{Position: position, Direction: EAST}, MaxX: 4, MaxY: 4}
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
				initialState = State{Robot: Robot{Position: position, Direction: NORTH}, MaxX: 4, MaxY: 4}
				finalState, _ = Reduce(action, initialState)
			})

			It("rotates the robot to EAST", func() {
				Expect(finalState.Robot.Direction).To(Equal(EAST))
			})
		})

		Context("when the robot is facing SOUTH", func() {
			BeforeEach(func() {
				initialState = State{Robot: Robot{Position: position, Direction: SOUTH}, MaxX: 4, MaxY: 4}
				finalState, _ = Reduce(action, initialState)
			})

			It("rotates the robot to EAST", func() {
				Expect(finalState.Robot.Direction).To(Equal(WEST))
			})
		})

		Context("when the robot is facing WEST", func() {
			BeforeEach(func() {
				initialState = State{Robot: Robot{Position: position, Direction: WEST}, MaxX: 4, MaxY: 4}
				finalState, _ = Reduce(action, initialState)
			})

			It("rotates the robot to SOUTH", func() {
				Expect(finalState.Robot.Direction).To(Equal(NORTH))
			})
		})

		Context("when the robot is facing EAST", func() {
			BeforeEach(func() {
				initialState = State{Robot: Robot{Position: position, Direction: EAST}, MaxX: 4, MaxY: 4}
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
