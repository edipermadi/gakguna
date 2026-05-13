package ethernet

type Frame struct {
	DestinationMAC [MacAddressLength]byte
	SourceMAC      [MacAddressLength]byte
	Tag            [TagLength]byte
	Type           [TypeLength]byte
	Payload        [PayloadLength]byte
}
