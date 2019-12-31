package elem

import (
	"bytes"
	"encoding/binary"
	"log"
)

type CreatePDR struct {
	EType   IEType
	ELength uint16
	PDI
	Precedence
	PDRID
	OuterHeaderRemoval
	FARID
	URRIDs []URRID
	QERIDs []QERID
	//TODO::
}

func DecodeCreatePDR(data []byte, len uint16) *CreatePDR {
	createPDR := CreatePDR{
		EType:   IETypeCreatedPDR,
		ELength: len,
		URRIDs:  []URRID{},
		QERIDs:  []QERID{},
	}
	var cursor uint16
	buf := bytes.NewBuffer(data)
	for cursor < createPDR.ELength {
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
		case IETypePDI:
			createPDR.PDI = *DecodePDI(eValue, eLen)
		case IETypePrecedence:
			createPDR.Precedence = *DecodePrecedence(eValue, eLen)
		case IETypePDRID:
			createPDR.PDRID = *DecodePDRID(eValue, eLen)
		case IETypeOuterHeaderRemoval:
			createPDR.OuterHeaderRemoval = *DecodeOuterHeaderRemoval(eValue, eLen)
		case IETypeFARID:
			createPDR.FARID = *DecodeFARID(eValue, eLen)
		case IETypeURRID:
			createPDR.URRIDs = append(createPDR.URRIDs, *DecodeURRID(eValue, eLen))
		case IETypeQERID:
			createPDR.QERIDs = append(createPDR.QERIDs, *DecodeQERID(eValue, eLen))
		default:
			log.Println("err: unknown tlv type", eType) //TODO::
		}
		cursor = cursor + eLen + 4
	}
	return &createPDR
}

func EncodeCreatePDR(createPDR CreatePDR) []byte {
	ret := setValue(createPDR.EType, createPDR.ELength, createPDR.PDI, createPDR.Precedence, createPDR.PDRID) //PDI Precedence PDRID 为M信元
	if HasOuterHeaderRemoval(createPDR.OuterHeaderRemoval) {
		ret = setValue(ret, createPDR.OuterHeaderRemoval)
	}
	if HasFARID(createPDR.FARID) {
		ret = setValue(ret, createPDR.FARID)
	}
	for urrID := range createPDR.URRIDs {
		ret = setValue(ret, urrID)
	}
	for qerID := range createPDR.QERIDs {
		ret = setValue(ret, qerID)
	}
	return ret
}

//判断是否含有CreatePDR
func Has(createPDR CreatePDR) bool {
	if createPDR.EType == 0 {
		return false
	}
	return true
}
