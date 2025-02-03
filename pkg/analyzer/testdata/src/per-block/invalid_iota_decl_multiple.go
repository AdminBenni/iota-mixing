package per_individual

const ( // want "iota mixing. keep iotas in separate blocks to consts with r-val"
	InvalidIotaDeclMultipleAbove   = "above"
	InvalidIotaDeclMultipleNotZero = iota
	InvalidIotaDeclMultipleNotOne
	InvalidIotaDeclMultipleBetween = "between"
	InvalidIotaDeclMultipleNotTwo
	InvalidIotaDeclMultipleBelow = "below"
)
