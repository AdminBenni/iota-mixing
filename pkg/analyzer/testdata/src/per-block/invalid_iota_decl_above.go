package per_block 

const ( // want "iota mixing. keep iotas in separate blocks to consts with r-val"
	InvalidIotaDeclAboveAnything = "anything"
	InvalidIotaDeclAboveNotZero  = iota
	InvalidIotaDeclAboveNotOne
	InvalidIotaDeclAboveNotTwo
)
