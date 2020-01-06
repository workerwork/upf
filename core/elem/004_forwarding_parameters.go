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

func DecodeForwardingParameters(buf *bytes.Buffer, len uint16) *ForwardingParameters {
	fps := ForwardingParameters{
		EType:   IETypeForwardingParameters,
		ELength: len,
	}
	var cursor uint16
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
		e := make([]byte, eLen)
		if err := binary.Read(buf, binary.BigEndian, &e); err != nil {
			log.Println(err) //TODO::
		}
		eValue := bytes.NewBuffer(e)
		switch eType {
		case IETypeTransportLevelMarking:
			fps.TransportLevelMarking = *DecodeTransportLevelMarking(eValue, eLen)
		case IETypeDestinationInterface:
			fps.DestinationInterface = *DecodeDestinationInterface(eValue, eLen)
		case IETypeNetworkInstance:
			fps.NetworkInstance = *DecodeNetworkInstance(eValue, eLen)
		case IETypeOuterHeaderCreation:
			fps.OuterHeaderCreation = *DecodeOuterHeaderCreation(eValue, eLen)
		default:
			log.Println("forwarding parameters err: unknown tlv type", eType) //TODO::
		}
		cursor = cursor + eLen + 4
	}
	return &fps
}

func EncodeForwardingParameters(fps ForwardingParameters) *bytes.Buffer {
	ret := SetValue(fps.EType, fps.ELength)
	switch {
	case HasDestinationInterface(fps.DestinationInterface): //M
		SetValue(ret, fps.DestinationInterface)
		fallthrough
	case HasNetworkInstance(fps.NetworkInstance):
		SetValue(ret, fps.NetworkInstance)
		fallthrough
	case HasTransportLevelMarking(fps.TransportLevelMarking):
		SetValue(ret, fps.TransportLevelMarking)
		fallthrough
	case HasOuterHeaderCreation(fps.OuterHeaderCreation):
		SetValue(ret, fps.OuterHeaderCreation)
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
