package per_individual

const (
	InvalidIotaDeclAboveAnything = "anything" // want "InvalidIotaDeclAboveAnything is a const with r-val in same const block as iota. keep iotas in separate const blocks"
	InvalidIotaDeclAboveNotZero  = iota
	InvalidIotaDeclAboveNotOne
	InvalidIotaDeclAboveNotTwo
)
