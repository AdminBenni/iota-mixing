package per_individual

const ( // want "iota mixing. keep iotas in separate blocks to consts with r-val"
	InvalidIotaDeclBetweenZero = iota
	InvalidIotaDeclBetweenOne
	InvalidIotaDeclBetweenAnything = "anything"
	InvalidIotaDeclBetweenNotTwo
)
