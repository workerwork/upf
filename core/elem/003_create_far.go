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

func DecodeCreateFAR(buf *bytes.Buffer, len uint16) *CreateFAR {
	createFAR := CreateFAR{
		EType:   IETypeCreateFAR,
		ELength: len,
	}
	var cursor uint16
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
		e := make([]byte, eLen)
		if err := binary.Read(buf, binary.BigEndian, &e); err != nil {
			log.Println(err) //TODO::
		}
		eValue := bytes.NewBuffer(e)
		switch eType {
		case IETypeForwardingParameters:
			createFAR.ForwardingParameters = *DecodeForwardingParameters(eValue, eLen)
		case IETypeApplyAction:
			createFAR.ApplyAction = *DecodeApplyAction(eValue, eLen)
		case IETypeFARID:
			createFAR.FARID = *DecodeFARID(eValue, eLen)
		default:
			log.Println(" create far err: unknown tlv type", eType) //TODO::
		}
		cursor = cursor + eLen + 4
	}
	return &createFAR
}

func EncodeCreateFAR(createFAR CreateFAR) []byte {
	ret := SetValue(createFAR.EType, createFAR.ELength)
	switch {
	case HasApplyAction(createFAR.ApplyAction): //M
		SetValue(ret, createFAR.ApplyAction)
		fallthrough
	case HasFARID(createFAR.FARID): //M
		SetValue(ret, createFAR.FARID)
		fallthrough
	case HasForwardingParameters(createFAR.ForwardingParameters):
		SetValue(ret, createFAR.ForwardingParameters)
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
