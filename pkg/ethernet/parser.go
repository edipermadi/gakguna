package ethernet

import "io"

type Parser struct {
	Buffer io.ByteReader
}

type stateCode int

const (
	stateWaitingForPreamble             stateCode = iota
	stateReceivingPreamble              stateCode = iota
	stateWaitingForSFD                  stateCode = iota
	stateReceivingDestinationMacAddress stateCode = iota
	stateReceivingSourceMacAddress      stateCode = iota
	stateReceivingTag                   stateCode = iota
	stateReceivingType                  stateCode = iota
	stateReceivingPayload               stateCode = iota
)

func (p *Parser) Parser() (*Frame, error) {

	var frame Frame
	state := stateWaitingForPreamble
	counter := 0
	for {
		b, err := p.Buffer.ReadByte()
		if err != nil {
			return nil, err
		}

		switch state {
		case stateWaitingForPreamble:
			if b == preambleByte {
				state = stateReceivingPreamble
				counter++
			}
		case stateReceivingPreamble:
			if b != preambleByte {
				return nil, ErrInvalidEthernetFramePreamble
			}
			counter++
			if counter == preambleBytesCount {
				counter = 0
				state = stateWaitingForSFD
			}
		case stateWaitingForSFD:
			if b != sfdByte {
				return nil, ErrInvalidEthernetFrameSFD
			}
			counter = 0
			state = stateReceivingDestinationMacAddress
		case stateReceivingDestinationMacAddress:
			frame.DestinationMAC[counter] = b
			counter++
			if counter == MacAddressLength {
				counter = 0
				state = stateReceivingSourceMacAddress
			}
		case stateReceivingSourceMacAddress:
			frame.SourceMAC[counter] = b
			counter++
			if counter == MacAddressLength {
				counter = 0
				state = stateReceivingSourceMacAddress
			}
		case stateReceivingTag:
			frame.Tag[counter] = b
			counter++
			if counter == TagLength {
				counter = 0
				state = stateReceivingType
			}
		case stateReceivingType:
			frame.Type[counter] = b
			counter++
			if counter == TypeLength {
				counter = 0
				state = stateReceivingPayload
			}
		case stateReceivingPayload:
			frame.Payload[counter] = b
			counter++
			if counter == PayloadLength {
				counter = 0
				state = stateReceivingPayload
			}
		}
	}

}
