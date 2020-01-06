package elem

import (
	"bytes"
	"encoding/binary"
	"log"
)

type CreateQER struct {
	EType   IEType
	ELength uint16
	GateStatus
	MBR
	QERID
	QFI
	//TODO::
}

func DecodeCreateQER(buf *bytes.Buffer, len uint16) *CreateQER {
	createQER := CreateQER{
		EType:   IETypeCreateQER,
		ELength: len,
	}
	var cursor uint16
	for cursor < createQER.ELength {
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
		case IETypeGateStatus:
			createQER.GateStatus = *DecodeGateStatus(eValue, eLen)
		case IETypeMBR:
			createQER.MBR = *DecodeMBR(eValue, eLen)
		case IETypeQERID:
			createQER.QERID = *DecodeQERID(eValue, eLen)
		case IETypeQFI:
			createQER.QFI = *DecodeQFI(eValue, eLen)
		default:
			log.Println("create qer err: unknown tlv type", eType) //TODO::
		}
		cursor = cursor + eLen + 4
	}
	return &createQER
}

func EncodeCreateQER(createQER CreateQER) []byte {
	ret := SetValue(createQER.EType, createQER.ELength)
	switch {
	case HasGateStatus(createQER.GateStatus): //M
		SetValue(ret, EncodeGateStatus(createQER.GateStatus))
		fallthrough
	case HasQERID(createQER.QERID): //M
		SetValue(ret, EncodeQERID(createQER.QERID))
		fallthrough
	case HasMBR(createQER.MBR):
		SetValue(ret, EncodeMBR(createQER.MBR))
		fallthrough
	case HasQFI(createQER.QFI):
		SetValue(ret, EncodeQFI(createQER.QFI))
	}
	return ret
}

//判断是否含有CreateQER
func HasCreateQER(createQER CreateQER) bool {
	if createQER.EType == 0 {
		return false
	}
	return true
}
