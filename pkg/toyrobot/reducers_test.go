package toyrobot

import (
	"../reporter"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestReducers(t *testing.T) {
	RegisterFailHandler(Fail)
	//RunSpecs(t, "Reducers Suite")
	RunSpecsWithCustomReporters(t, "Toy Robot Reducer", []Reporter{reporter.New()})
}

var _ = Describe("Reducer", func() {
	var action Action
	var initialState State
	var finalState State
	var err error

	BeforeEach(func() {
		initialState = State{Robot: Robot{2, 3, SOUTH}}
	})

	Describe("reduces LEFT action", func() {
		BeforeEach(func() {
			action = Action{ActionType: LEFT}
		})

		Context("when the robot is facing NORTH", func() {
			BeforeEach(func() {
				initialState = State{Robot: Robot{2, 3, NORTH}}
				finalState, _ = Reduce(action, initialState)
			})

			It("rotates the robot to EAST", func() {
				Expect(finalState.Robot.Direction).To(Equal(WEST))
			})
		})

		Context("when the robot is facing SOUTH", func() {
			BeforeEach(func() {
				initialState = State{Robot: Robot{2, 3, SOUTH}}
				finalState, _ = Reduce(action, initialState)
			})

			It("rotates the robot to EAST", func() {
				Expect(finalState.Robot.Direction).To(Equal(EAST))
			})
		})

		Context("when the robot is facing WEST", func() {
			BeforeEach(func() {
				initialState = State{Robot: Robot{2, 3, WEST}}
				finalState, _ = Reduce(action, initialState)
			})

			It("rotates the robot to SOUTH", func() {
				Expect(finalState.Robot.Direction).To(Equal(SOUTH))
			})
		})

		Context("when the robot is facing EAST", func() {
			BeforeEach(func() {
				initialState = State{Robot: Robot{2, 3, EAST}}
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
				initialState = State{Robot: Robot{2, 3, NORTH}}
				finalState, _ = Reduce(action, initialState)
			})

			It("rotates the robot to EAST", func() {
				Expect(finalState.Robot.Direction).To(Equal(EAST))
			})
		})

		Context("when the robot is facing SOUTH", func() {
			BeforeEach(func() {
				initialState = State{Robot: Robot{2, 3, SOUTH}}
				finalState, _ = Reduce(action, initialState)
			})

			It("rotates the robot to EAST", func() {
				Expect(finalState.Robot.Direction).To(Equal(WEST))
			})
		})

		Context("when the robot is facing WEST", func() {
			BeforeEach(func() {
				initialState = State{Robot: Robot{2, 3, WEST}}
				finalState, _ = Reduce(action, initialState)
			})

			It("rotates the robot to SOUTH", func() {
				Expect(finalState.Robot.Direction).To(Equal(NORTH))
			})
		})

		Context("when the robot is facing EAST", func() {
			BeforeEach(func() {
				initialState = State{Robot: Robot{2, 3, EAST}}
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
