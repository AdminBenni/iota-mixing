package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"

	"github.com/AdminBenni/iota-mixing/pkg/analyzer"
	"github.com/AdminBenni/iota-mixing/pkg/analyzer/flags"
)

func main() {
	iotaMixingAnalyzer := analyzer.GetIotaMixingAnalyzer()

	flags.SetupFlags(&iotaMixingAnalyzer.Flags)

	singlechecker.Main(iotaMixingAnalyzer)
}
