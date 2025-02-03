package flags

import "flag"

const (
	TrueString  = "true"
	FalseString = "false"

	ReportIndividualFlagName  = "report-individual"
	reportIndividualFlagUsage = "whether or not to report individual consts rather than just the const block."
)

var (
	ReportIndividualFlag *string
)

func SetupFlags(flags *flag.FlagSet) {
	ReportIndividualFlag = flags.String(ReportIndividualFlagName, FalseString, reportIndividualFlagUsage)
}
