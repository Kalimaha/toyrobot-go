package reporter

import (
	"fmt"
	"github.com/onsi/ginkgo/config"
	"github.com/onsi/ginkgo/types"
)

type rspecReporter struct {
	leaves []Leaf
}

type Leaf struct {
	Description string
	Leaves      []Leaf
	Passed      int
}

func New() *rspecReporter {
	return &rspecReporter{
		leaves: []Leaf{},
	}
}

func (r *rspecReporter) SpecSuiteWillBegin(config config.GinkgoConfigType, summary *types.SuiteSummary) {}
func (r *rspecReporter) BeforeSuiteDidRun(setupSummary *types.SetupSummary) {}
func (r *rspecReporter) AfterSuiteDidRun(setupSummary *types.SetupSummary) {}

func (r *rspecReporter) SpecWillRun(spec *types.SpecSummary) {}

func printLeaves(leaves []Leaf) {
	for _, v := range leaves {
		fmt.Println(v.Description, icon(v.Passed))
		for _, v2 := range v.Leaves {
			fmt.Println(v2.Description, icon(v2.Passed))
		}
	}
}

func icon(passed int) string {
	if passed == 1 {
		return "✔️ "
	} else if passed == -1 {
		return "x "
	}
	return ""
}

func bool2int(passed bool) int {
	if passed == true {
		return 1
	}
	return -1
}

func (r *rspecReporter) SpecDidComplete(spec *types.SpecSummary) {
	fmt.Println(spec.ComponentTexts)
	foundLevelOne := false
	for idx, v := range r.leaves {
		if v.Description == fmt.Sprintf("  %s", spec.ComponentTexts[1]) {
			foundLevelOne = true
			newLeaf := Leaf{Description: fmt.Sprintf("    %s", spec.ComponentTexts[2]), Passed: bool2int(spec.Passed())}
			r.leaves[idx].Leaves = append(r.leaves[idx].Leaves, newLeaf)

			foundLevelTwo := false
			for idx2, v2 := range r.leaves {
				if v.Description == fmt.Sprintf("  %s", spec.ComponentTexts[1]) && v2.Description == fmt.Sprintf("    %s", spec.ComponentTexts[2]) {
					foundLevelTwo = true
					fmt.Println("WE FOUND", v2)

					newNewLeaf := Leaf{Description: fmt.Sprintf("      %s", spec.ComponentTexts[3]), Passed: bool2int(spec.Passed())}
					r.leaves[idx2].Leaves = append(r.leaves[idx2].Leaves, newNewLeaf)
				}
			}
			if foundLevelTwo == false {
				v.Leaves = append(v.Leaves, Leaf{Description: fmt.Sprintf("    %s", spec.ComponentTexts[2]), Passed: bool2int(spec.Passed())})
			}
		}
	}

	if foundLevelOne == false {
		newLeaf := Leaf{
			Description: fmt.Sprintf("  %s", spec.ComponentTexts[1]),
			Leaves: []Leaf{{Description: fmt.Sprintf("    %s", spec.ComponentTexts[2]), Passed: bool2int(spec.Passed())}},
		}
		r.leaves = append(r.leaves, newLeaf)
	}
}

func (r *rspecReporter) SpecSuiteDidEnd(summary *types.SuiteSummary) {
	fmt.Println()
	fmt.Println(summary.SuiteDescription)
	//fmt.Println(summary.SuiteDescription, "outcome:")
	printLeaves(r.leaves)
	fmt.Println()
	//fmt.Println(summary.SuiteDescription, "outcome:")
}
