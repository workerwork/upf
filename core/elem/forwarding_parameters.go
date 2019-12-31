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
	fps := ForwardingParameters{
		EType:   IETypePDI,
		ELength: len,
	}
	var cursor uint16
	buf := bytes.NewBuffer(data)
	for cursor < fps.ELength {
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
			fps.TransportLevelMarking = *DecodeTransportLevelMarking(eValue, eLen)
		case IETypeFTEID:
			fps.DestinationInterface = *DecodeDestinationInterface(eValue, eLen)
		case IETypeNetworkInstance:
			fps.NetworkInstance = *DecodeNetworkInstance(eValue, eLen)
		case IETypeQFI:
			fps.OuterHeaderCreation = *DecodeOuterHeaderCreation(eValue, eLen)
		default:
			log.Println("err: unknown tlv type", eType) //TODO::
		}
		cursor = cursor + eLen + 4
	}
	return &fps
}

func EncodeForwardingParameters(fps ForwardingParameters) []byte {
	ret := setValue(fps.EType, fps.ELength, fps.DestinationInterface) //DestinationInterface 为M信元
	if HasNetworkInstance(fps.NetworkInstance) {
		ret = setValue(ret, fps.NetworkInstance)
	}
	if HasTransportLevelMarking(fps.TransportLevelMarking) {
		ret = setValue(ret, fps.TransportLevelMarking)
	}
	if HasOuterHeaderCreation(fps.OuterHeaderCreation) {
		ret = setValue(ret, fps.OuterHeaderCreation)
	}
	return ret
}

//判断是否含有ForwardingParameters
func HasForwardingParameters(fps ForwardingParameters) bool {
	if fps.EType == 0 {
		return false
	}
	return true
}

