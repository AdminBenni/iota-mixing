package per_individual

const (
	InvalidIotaDeclBelowZero = iota
	InvalidIotaDeclBelowOne
	InvalidIotaDeclBelowTwo
	InvalidIotaDeclBelowAnything = "anything" // want "InvalidIotaDeclBelowAnything is a const with r-val in same const block as iota. keep iotas in separate const blocks"
)
