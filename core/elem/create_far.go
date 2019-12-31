package elem

import (
	"bytes"
	"encoding/binary"
	"log"
)

type CreateFAR struct {
	EType   IEType
	ELength uint16
	ForwardingParameters
	ApplyAction
	FARID
	//TODO::
}

func DecodeCreateFAR(data []byte, len uint16) *CreateFAR {
	createFAR := CreateFAR{
		EType:   IETypeCreateFAR,
		ELength: len,
	}
	var cursor uint16
	buf := bytes.NewBuffer(data)
	for cursor < createFAR.ELength {
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
		case IETypeForwardingParameters:
			createFAR.ForwardingParameters = *DecodeForwardingParameters(eValue, eLen)
		case IETypeApplyAction:
			createFAR.ApplyAction = *DecodeApplyAction(eValue, eLen)
		case IETypeFARID:
			createFAR.FARID = *DecodeFARID(eValue, eLen)
		default:
			log.Println("err: unknown tlv type", eType) //TODO::
		}
		cursor = cursor + eLen + 4
	}
	return &createFAR
}

func EncodeCreateFAR(createFAR CreateFAR) []byte {
	ret := setValue(createFAR.EType, createFAR.ELength, createFAR.ApplyAction, createFAR.FARID) //ApplyAction FARID 为M信元
	if HasFARID(createFAR.FARID) {
		ret = setValue(ret, createFAR.FARID)
	}
	return ret
}

//判断是否含有CreateFAR
func HasCreateFAR(createFAR CreateFAR) bool {
	if createFAR.EType == 0 {
		return false
	}
	return true
}
