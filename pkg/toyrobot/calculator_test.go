package calculator

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestCalculator(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Calculator Suite")
}

var _ = Describe("Sum", func() {
	var a int
	var b int

	BeforeEach(func() {
		a = 2
		b = 3
	})

	It("sums two numbers", func() {
		Expect(Sum(a, b)).To(Equal(5))
	})
})
