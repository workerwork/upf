package elem

import (
	"bytes"
	"encoding/binary"
	"log"
)

type CreateURR struct {
	EType   IEType
	ELength uint16
	ReportingTriggers
	MeasurementMethod
	MeasurementPeriod
	URRID
	//TODO::
}

func DecodeCreateURR(data []byte, len uint16) *CreateURR {
	createURR := CreateURR{
		EType:   IETypeCreateURR,
		ELength: len,
	}
	var cursor uint16
	buf := bytes.NewBuffer(data)
	for cursor < createURR.ELength {
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
			createURR.ReportingTriggers = *DecodeReportingTriggers(eValue, eLen)
		case IETypeApplyAction:
			createURR.MeasurementMethod = *DecodeMeasurementMethod(eValue, eLen)
		case IETypeFARID:
			createURR.MeasurementPeriod = *DecodeMeasurementPeriod(eValue, eLen)
		default:
			log.Println("err: unknown tlv type", eType) //TODO::
		}
		cursor = cursor + eLen + 4
	}
	return &createURR
}

func EncodeCreateURR(createURR CreateURR) []byte {
	ret := setValue(createURR.EType, createURR.ELength, createURR.ReportingTriggers, createURR.MeasurementMethod, createURR.URRID) //ReportingTriggers MeasurementMethod URRID为M信元
	if HasMeasurementPeriod(createURR.MeasurementPeriod) {
		ret = setValue(ret, createURR.MeasurementPeriod)
	}
	return ret
}

//判断是否含有CreateURR
func HasCreateURR(createURR CreateURR) bool {
	if createURR.EType == 0 {
		return false
	}
	return true
}
