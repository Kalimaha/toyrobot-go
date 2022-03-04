package toyrobot

import (
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"os"
	"testing"
)

func TestParser(t *testing.T) {
	RegisterFailHandler(Fail)
}

var _ = Describe("Parse File", func() {
	var actions []Action
	var err error
	var currentPath string
	var filepath string

	BeforeEach(func() {
		currentPath, _ = os.Getwd()
		filepath = fmt.Sprintf("%s/../../resources/testParseOne.txt", currentPath)
		actions, err = ParseFile(filepath)
	})

	It("returns a list of actions", func() {
		Expect(len(actions)).To(Equal(5))
	})

	Context("when the file does NOT exist", func() {
		BeforeEach(func() {
			filepath = fmt.Sprintf("%s/../../resources/spam.txt", currentPath)
			actions, err = ParseFile(filepath)
		})

		It("returns an error", func() {
			Expect(err.Error()).To(Equal("file does not exist"))
		})
	})

	Context("when the file contains and invalid action", func() {
		BeforeEach(func() {
			filepath = fmt.Sprintf("%s/../../resources/testParseTwo.txt", currentPath)
			actions, err = ParseFile(filepath)
		})

		It("discards the invalid action", func() {
			Expect(len(actions)).To(Equal(5))
		})
	})

	Context("when the file doesn't start with a PLACE action", func() {
		BeforeEach(func() {
			filepath = fmt.Sprintf("%s/../../resources/testParseThree.txt", currentPath)
			actions, err = ParseFile(filepath)
		})

		It("discards the invalid actions", func() {
			Expect(len(actions)).To(Equal(5))
		})
	})
})

var _ = Describe("Parse Line", func() {
	It("parses 'MAP' string to a MAP action", func() {
		Expect(ParseLine("MAP")).To(Equal(Action{ActionType: MAP}))
	})

	It("parses 'PLACE_OBJECT' string to a PLACE_OBJECT action", func() {
		Expect(ParseLine("PLACE_OBJECT")).To(Equal(Action{ActionType: PLACE_OBJECT}))
	})

	It("parses 'MOVE' string to a MOVE action", func() {
		Expect(ParseLine("MOVE")).To(Equal(Action{ActionType: MOVE}))
	})

	It("parses 'LEFT' string to a LEFT action", func() {
		Expect(ParseLine("LEFT")).To(Equal(Action{ActionType: LEFT}))
	})

	It("parses 'RIGHT' string to a RIGHT action", func() {
		Expect(ParseLine("RIGHT")).To(Equal(Action{ActionType: RIGHT}))
	})

	It("parses 'REPORT' string to a REPORT action", func() {
		Expect(ParseLine("REPORT")).To(Equal(Action{ActionType: REPORT}))
	})

	It("parses 'PLACE 2,3,NORTH' string to a PLACE action", func() {
		Expect(ParseLine("PLACE 2,3,NORTH")).To(Equal(Action{ActionType: PLACE, X: 2, Y: 3, Direction: NORTH}))
	})

	It("parses 'PLACE 2,3,SOUTH' string to a PLACE action", func() {
		Expect(ParseLine("PLACE 2,3,SOUTH")).To(Equal(Action{ActionType: PLACE, X: 2, Y: 3, Direction: SOUTH}))
	})

	It("parses 'PLACE 2,3,EAST' string to a PLACE action", func() {
		Expect(ParseLine("PLACE 2,3,EAST")).To(Equal(Action{ActionType: PLACE, X: 2, Y: 3, Direction: EAST}))
	})

	It("parses 'PLACE 2,3,WEST' string to a PLACE action", func() {
		Expect(ParseLine("PLACE 2,3,WEST")).To(Equal(Action{ActionType: PLACE, X: 2, Y: 3, Direction: WEST}))
	})

	Context("when the action is not valid", func() {
		It("returns an error", func() {
			_, err := ParseLine("SPAM")
			Expect(err.Error()).To(Equal("invalid action: SPAM"))
		})
	})
})
