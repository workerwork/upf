package elem

import (
	"bytes"
	"encoding/binary"
	"log"
)

type ForwardingParameters struct {
	EType   IEType
	ELength uint16
	NetworkInstance
	TransportLevelMarking
	DestinationInterface
	OuterHeaderCreation
	//TODO::
}

func DecodeForwardingParameters(data []byte, len uint16) *ForwardingParameters {
	forwardingParameters := ForwardingParameters{
		EType:   IETypePDI,
		ELength: len,
	}
	var cursor uint16
	buf := bytes.NewBuffer(data)
	for cursor < forwardingParameters.ELength {
		var (
			eType IEType
			eLen  uint16
		)
		if err := binary.Read(buf, binary.BigEndian, &eType); err != nil {
			log.Println(err) //TODO::
		}
		if err := binary.Read(buf, binary.BigEndian, &eLen); err != nil {
			log.Println(err) //TODO::
		}
		eValue := make([]byte, eLen)
		if err := binary.Read(buf, binary.BigEndian, &eValue); err != nil {
			log.Println(err) //TODO::
		}
		switch eType {
		case IETypeSourceInterface:
			forwardingParameters.TransportLevelMarking = *DecodeTransportLevelMarking(eValue, eLen)
		case IETypeFTEID:
			forwardingParameters.DestinationInterface = *DecodeDestinationInterface(eValue, eLen)
		case IETypeNetworkInstance:
			forwardingParameters.NetworkInstance = *DecodeNetworkInstance(eValue, eLen)
		case IETypeQFI:
			forwardingParameters.OuterHeaderCreation = *DecodeOuterHeaderCreation(eValue, eLen)
		default:
			log.Println("err: unknown tlv type", eType) //TODO::
		}
		cursor = cursor + eLen + 4
	}
	return &forwardingParameters
}

func EncodeForwardingParameters(forwardingParameters ForwardingParameters) []byte {
	ret := setValue(forwardingParameters.EType, forwardingParameters.ELength, forwardingParameters.DestinationInterface) //DestinationInterface 为M信元
	if HasNetworkInstance(forwardingParameters.NetworkInstance) {
		ret = setValue(ret, forwardingParameters.NetworkInstance)
	}
	if HasTransportLevelMarking(forwardingParameters.TransportLevelMarking) {
		ret = setValue(ret, forwardingParameters.TransportLevelMarking)
	}
	if HasOuterHeaderCreation(forwardingParameters.OuterHeaderCreation) {
		ret = setValue(ret, forwardingParameters.OuterHeaderCreation)
	}
	return ret
}

//判断是否含有ForwardingParameters
func HasForwardingParameters(forwardingParameters ForwardingParameters) bool {
	if forwardingParameters.EType == 0 {
		return false
	}
	return true
}

