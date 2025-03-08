package analyzer_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/AdminBenni/iota-mixing/pkg/analyzer"
	"github.com/AdminBenni/iota-mixing/pkg/analyzer/flags"
)

func TestAnalyzerPerBlock(t *testing.T) {
	iotaMixingAnalyzer := analyzer.GetIotaMixingAnalyzer()

	flags.SetupFlags(&iotaMixingAnalyzer.Flags)

	err := iotaMixingAnalyzer.Flags.Set(flags.ReportIndividualFlagName, flags.FalseString)
	if err != nil {
		t.Fatal(err)
	}

	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, iotaMixingAnalyzer, "per-block")
}

func TestAnalyzerPerIndividual(t *testing.T) {
	iotaMixingAnalyzer := analyzer.GetIotaMixingAnalyzer()

	flags.SetupFlags(&iotaMixingAnalyzer.Flags)

	err := iotaMixingAnalyzer.Flags.Set(flags.ReportIndividualFlagName, flags.TrueString)
	if err != nil {
		t.Fatal(err)
	}

	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, iotaMixingAnalyzer, "per-individual")
}
