package golandreporter

import (
	"fmt"
	"os"
	"strings"

	"github.com/onsi/ginkgo/config"
	"github.com/onsi/ginkgo/reporters"
	"github.com/onsi/ginkgo/reporters/stenographer"
	"github.com/onsi/ginkgo/types"
)

type GolandReporter struct{}

func NewGolandReporter() reporters.Reporter {
	return GolandReporter{}
}

func NewAutoGolandReporter() reporters.Reporter {
	if strings.Contains(os.Getenv("OLDPWD"), "Goland") {
		return NewGolandReporter()
	} else {
		stenographer := stenographer.New(!config.DefaultReporterConfig.NoColor, config.GinkgoConfig.FlakeAttempts > 1, os.Stdout)
		return reporters.NewDefaultReporter(config.DefaultReporterConfig, stenographer)
	}
}

func (g GolandReporter) SpecSuiteWillBegin(config config.GinkgoConfigType, summary *types.SuiteSummary) {

}

func (g GolandReporter) BeforeSuiteDidRun(setupSummary *types.SetupSummary) {

}

func (g GolandReporter) SpecWillRun(specSummary *types.SpecSummary) {
	fmt.Printf("=== RUN   %s\n", strings.Join(specSummary.ComponentTexts[1:], " "))
}

func (g GolandReporter) SpecDidComplete(specSummary *types.SpecSummary) {
	if specSummary.Passed() {
		printResultOutput(specSummary, "PASS")
	} else if specSummary.HasFailureState() {
		fmt.Printf("%s\n\n", specSummary.Failure.Message)
		fmt.Printf("%s\n\n", specSummary.Failure.Location.FullStackTrace)
		printResultOutput(specSummary, "FAIL")
	} else if specSummary.Skipped() {
		printResultOutput(specSummary, "SKIP")
	} else if specSummary.Pending() {
		printResultOutput(specSummary, "SKIP")
	} else {
		panic("Unknown test output")
	}
}

func (g GolandReporter) AfterSuiteDidRun(setupSummary *types.SetupSummary) {

}

func (g GolandReporter) SpecSuiteDidEnd(summary *types.SuiteSummary) {

}

func printResultOutput(specSummary *types.SpecSummary, result string) {
	fmt.Printf("--- %s: %s (%.3fs)\n", result, strings.Join(specSummary.ComponentTexts[1:], " "), specSummary.RunTime.Seconds())
}
