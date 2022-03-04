package toyrobot

import (
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"os"
	"testing"
)

func TestPlay(t *testing.T) {
	RegisterFailHandler(Fail)
}

var _ = Describe("Play", func() {
	var currentPath string
	var filepath string
	var state State
	var err error

	BeforeEach(func() {
		currentPath, _ = os.Getwd()
	})

	It("runs the simulation for Example A", func() {
		filepath = fmt.Sprintf("%s/../../resources/exampleA.txt", currentPath)
		state, _ = Play(filepath, ERROR)

		Expect(state.Report).To(Equal("The robot is at (0, 1) and it is facing North."))
	})

	It("runs the simulation for Example B", func() {
		filepath = fmt.Sprintf("%s/../../resources/exampleB.txt", currentPath)
		state, _ = Play(filepath, ERROR)

		Expect(state.Report).To(Equal("The robot is at (0, 0) and it is facing West."))
	})

	It("runs the simulation for Example C", func() {
		filepath = fmt.Sprintf("%s/../../resources/exampleC.txt", currentPath)
		state, _ = Play(filepath, ERROR)

		Expect(state.Report).To(Equal("The robot is at (3, 3) and it is facing North."))
	})

	It("runs the simulation for Example D", func() {
		filepath = fmt.Sprintf("%s/../../resources/exampleD.txt", currentPath)
		state, _ = Play(filepath, ERROR)

		Expect(state.Report).To(Equal("The robot is at (3, 3) and it is facing North."))
	})

	Context("when the file does not exist", func() {
		BeforeEach(func() {
			_, err = Play(fmt.Sprintf("%s/../../resources/spam.txt", currentPath), ERROR)
		})

		It("returns an error", func() {
			Expect(err.Error()).To(Equal("file does not exist"))
		})
	})
})
