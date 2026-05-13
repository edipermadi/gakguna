package ethernet

import "errors"

var (
	ErrInvalidEthernetFramePreamble = errors.New("invalid Ethernet frame preamble")
	ErrInvalidEthernetFrameSFD      = errors.New("invalid Ethernet frame SFD")
)
