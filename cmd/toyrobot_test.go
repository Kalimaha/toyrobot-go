package main

import (
	"github.com/Kalimaha/ginkgo/reporter"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestInstructions(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecsWithCustomReporters(t, "Toy Robot - CLI", []Reporter{reporter.New()})
}

var _ = Describe("Instructions", func() {
	It("generates instructions for the user", func() {
		s := `
Please provide the absolute path to the file containing the instruction for the Toy Robot
e.g. go run ./toyrobot /tmp/example.txt
`
		Expect(Instructions()).To(Equal(s))
	})
})
