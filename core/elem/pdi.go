package elem

import (
	"bytes"
	"encoding/binary"
	"log"
)

type PDI struct {
	EType IEType
	ELength uint16
	SourceInterface
	FTEID
	NetworkInstance
	QFI
}

func DecodePDI(data []byte, len uint16) *PDI {
	var pdi PDI
	pdi.EType = IETypePDI
	pdi.ELength = len
	var cursor uint16
	buf := bytes.NewBuffer(data)
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
		eValue := make([]byte, eLen)
		if err := binary.Read(buf, binary.BigEndian, &eValue); err != nil {
			log.Println(err) //TODO::
		}
		switch eType {
		case IETypeSourceInterface:
			pdi.SourceInterface = *DecodeSourceInterface(eValue, eLen)
		case IETypeFTEID:
			pdi.FTEID = *DecodeFTEID(eValue, eLen)
		case IETypeNetworkInstance:
			pdi.NetworkInstance = *DecodeNetworkInstance(eValue,eLen)
		case IETypeQFI:
			pdi.QFI = *DecodeQFI(eValue,eLen)
		default:
			log.Println("err: unknown tlv type", eType)	//TODO::
		}
		cursor = cursor + eLen + 4
	}
	return &pdi
}

func EncodePDI(pdi PDI) []byte {
	ret := setValue(pdi.EType, pdi.ELength)
	if HasSourceInterface(pdi.SourceInterface) {
		ret = setValue(ret,pdi.SourceInterface)
	}
	if HasFTEID(pdi.FTEID) {
		ret = setValue(ret,pdi.FTEID)
	}
	if HasNetworkInstance(pdi.NetworkInstance) {
		ret = setValue(ret,pdi.NetworkInstance)
	}
	if HasQFI(pdi.QFI) {
		ret = setValue(ret,pdi.QFI)
	}
	return ret
}

//判断是否含有SourceInterface
func HasPDI(pdi PDI) bool {
	if pdi.EType == 0 {
		return false
	}
	return true
}