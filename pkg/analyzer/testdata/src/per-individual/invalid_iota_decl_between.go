package per_individual

const (
	InvalidIotaDeclBetweenZero = iota
	InvalidIotaDeclBetweenOne
	InvalidIotaDeclBetweenAnything = "anything" // want "InvalidIotaDeclBetweenAnything is a const with r-val in same const block as iota. keep iotas in separate const blocks"
	InvalidIotaDeclBetweenNotTwo
)
