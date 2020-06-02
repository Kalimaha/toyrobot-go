package toyrobot

import (
	"fmt"
	"github.com/onsi/ginkgo/config"
	"github.com/onsi/ginkgo/types"
	"strings"
)

type toyRobotReporter struct {
	// keep track of nested describes using map
	specs map[int]string
}

func New() *toyRobotReporter {
	return &toyRobotReporter{
		specs: make(map[int]string),
	}
}

func (r *toyRobotReporter) SpecSuiteWillBegin(config config.GinkgoConfigType, summary *types.SuiteSummary) {
}

func (r *toyRobotReporter) BeforeSuiteDidRun(setupSummary *types.SetupSummary) {
}

func (r *toyRobotReporter) AfterSuiteDidRun(setupSummary *types.SetupSummary) {
}

func (r *toyRobotReporter) SpecWillRun(spec *types.SpecSummary) {

}

func (r *toyRobotReporter) SpecDidComplete(spec *types.SpecSummary) {
	switch spec.State {
	case types.SpecStatePassed:
		fmt.Println("\n<PASSED::>Test Passed")
	case types.SpecStateFailed:
		fmt.Println("\n<FAILED::>Test Failed")
		fmt.Printf("\n<LOG:ESC:>%s\n", escape(spec.Failure.Message))
	case types.SpecStatePanicked:
		fmt.Printf("\n<ERROR::>%s\n", escape(spec.Failure.Message))
		fmt.Printf("\n<LOG::Panic>%s\n", escape(spec.Failure.ForwardedPanic))
		fmt.Printf("\n<TAB::Stack Trace>%s\n", escape(spec.Failure.Location.FullStackTrace))
	case types.SpecStateTimedOut:
	case types.SpecStateSkipped:
	case types.SpecStatePending:
	}
	fmt.Printf("\n<COMPLETEDIN::>%.4f\n", spec.RunTime.Seconds()*1000)
}

func (r *toyRobotReporter) SpecSuiteDidEnd(summary *types.SuiteSummary) {
	m := len(r.specs)
	for j := 1; j <= m; j++ {
		fmt.Println("\n<QWEQEWEWWE::>")
	}
}

func escape(s string) string {
	return strings.Replace(s, "\n", "<:LF:>", -1)
}
