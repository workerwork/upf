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
			log.Println("pdi.SourceInterface: ", pdi.SourceInterface)
		case IETypeFTEID:
			pdi.FTEID = *DecodeFTEID(eValue, eLen)
			log.Println("pdi.FTEID: ", pdi.FTEID)
		case IETypeNetworkInstance:
			pdi.NetworkInstance = *DecodeNetworkInstance(eValue, eLen)
			log.Println("pdi.NetworkInstance: ", pdi.NetworkInstance)
		case IETypeQFI:
			pdi.QFI = *DecodeQFI(eValue, eLen)
			log.Println("pdi.QFI: ", pdi.QFI)
		default:
			log.Println("pdi err: unknown tlv type", eType) //TODO::
		}
		cursor = cursor + eLen + 4
	}
	return &pdi
}

func EncodePDI(pdi PDI) []byte {
	ret := setValue(pdi.EType, pdi.ELength, pdi.SourceInterface) //SourceInterface 为M信元
	if HasFTEID(pdi.FTEID) {
		ret = setValue(ret, pdi.FTEID)
	}
	if HasNetworkInstance(pdi.NetworkInstance) {
		ret = setValue(ret, pdi.NetworkInstance)
	}
	if HasQFI(pdi.QFI) {
		ret = setValue(ret, pdi.QFI)
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
