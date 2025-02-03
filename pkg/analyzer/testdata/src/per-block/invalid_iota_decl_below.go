package per_individual

const ( // want "iota mixing. keep iotas in separate blocks to consts with r-val"
	InvalidIotaDeclBelowZero = iota
	InvalidIotaDeclBelowOne
	InvalidIotaDeclBelowTwo
	InvalidIotaDeclBelowAnything = "anything"
)
