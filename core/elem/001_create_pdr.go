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

func DecodeCreatePDR(buf *bytes.Buffer, len uint16) *CreatePDR {
	createPDR := CreatePDR{
		EType:   IETypeCreatePDR,
		ELength: len,
		URRIDs:  []URRID{},
		QERIDs:  []QERID{},
	}
	var cursor uint16
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
		e := make([]byte, eLen)
		if err := binary.Read(buf, binary.BigEndian, &e); err != nil {
			log.Println(err) //TODO::
		}
		eValue := bytes.NewBuffer(e)
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
			log.Println("create pdr err: unknown tlv type", eType) //TODO::
		}
		cursor = cursor + eLen + 4
	}
	return &createPDR
}

func EncodeCreatePDR(createPDR CreatePDR) *bytes.Buffer {
	ret := SetValue(createPDR.EType, createPDR.ELength)
	switch {
	case HasPDI(createPDR.PDI): //M
		SetValue(ret, createPDR.PDI)
		fallthrough
	case HasPrecedence(createPDR.Precedence): //M
		SetValue(ret, createPDR.Precedence)
		fallthrough
	case HasPDRID(createPDR.PDRID): //M
		SetValue(ret, createPDR.PDRID)
		fallthrough
	case HasOuterHeaderRemoval(createPDR.OuterHeaderRemoval):
		ret = SetValue(ret, createPDR.OuterHeaderRemoval)
		fallthrough
	case HasFARID(createPDR.FARID):
		ret = SetValue(ret, createPDR.FARID)
	}
	for _, urrID := range createPDR.URRIDs {
		if HasURRID(urrID) {
			ret = SetValue(ret, urrID)
		}
	}
	for _, qerID := range createPDR.QERIDs {
		if HasQERID(qerID) {
			ret = SetValue(ret, qerID)
		}
	}
	return ret
}

//判断是否含有CreatePDR
func HasCreatePDR(createPDR CreatePDR) bool {
	if createPDR.EType == 0 {
		return false
	}
	return true
}
