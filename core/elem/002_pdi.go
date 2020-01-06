package elem

import (
	"bytes"
	"encoding/binary"
	"log"
)

type PDI struct {
	EType   IEType
	ELength uint16
	SourceInterface
	FTEID
	NetworkInstance
	QFI
	UEIPAddress
}

func DecodePDI(buf *bytes.Buffer, len uint16) *PDI {
	pdi := PDI{
		EType:   IETypePDI,
		ELength: len,
	}
	var cursor uint16
	for cursor < pdi.ELength {
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
		case IETypeSourceInterface:
			pdi.SourceInterface = *DecodeSourceInterface(eValue, eLen)
		case IETypeFTEID:
			pdi.FTEID = *DecodeFTEID(eValue, eLen)
		case IETypeNetworkInstance:
			pdi.NetworkInstance = *DecodeNetworkInstance(eValue, eLen)
		case IETypeQFI:
			pdi.QFI = *DecodeQFI(eValue, eLen)
		case IETypeUEIPAddress:
			pdi.UEIPAddress = *DecodeUEIPAddress(eValue, eLen)
		default:
			log.Println("pdi err: unknown tlv type", eType) //TODO::
		}
		cursor = cursor + eLen + 4
	}
	//log.Println("pdi: ", pdi)
	return &pdi
}

func EncodePDI(pdi PDI) *bytes.Buffer {
	ret := SetValue(pdi.EType, pdi.ELength)
	switch {
	case HasSourceInterface(pdi.SourceInterface): //M
		SetValue(ret, pdi.SourceInterface)
		fallthrough
	case HasFTEID(pdi.FTEID):
		SetValue(ret, pdi.FTEID)
		fallthrough
	case HasNetworkInstance(pdi.NetworkInstance):
		SetValue(ret, pdi.NetworkInstance)
		fallthrough
	case HasQFI(pdi.QFI):
		SetValue(ret, pdi.QFI)
	}
	return ret
}

//判断是否含有PDI
func HasPDI(pdi PDI) bool {
	if pdi.EType == 0 {
		return false
	}
	return true
}
