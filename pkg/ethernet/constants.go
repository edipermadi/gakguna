package ethernet

const (
	preambleBytesCount = 7
	preambleByte       = byte(0x55)
	sfdByte            = byte(0xd5)
	MacAddressLength   = 6
	TagLength          = 4
	TypeLength         = 2
	PayloadLength      = 1500
)
