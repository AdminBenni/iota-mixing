package per_individual

const (
	InvalidIotaDeclMultipleAbove   = "above" // want "InvalidIotaDeclMultipleAbove is a const with r-val in same const block as iota. keep iotas in separate const blocks"
	InvalidIotaDeclMultipleNotZero = iota
	InvalidIotaDeclMultipleNotOne
	InvalidIotaDeclMultipleBetween = "between" // want "InvalidIotaDeclMultipleBetween is a const with r-val in same const block as iota. keep iotas in separate const blocks"
	InvalidIotaDeclMultipleNotTwo
	InvalidIotaDeclMultipleBelow = "below" // want "InvalidIotaDeclMultipleBelow is a const with r-val in same const block as iota. keep iotas in separate const blocks"
)
