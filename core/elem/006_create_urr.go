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

func DecodeCreateURR(buf *bytes.Buffer, len uint16) *CreateURR {
	createURR := CreateURR{
		EType:   IETypeCreateURR,
		ELength: len,
	}
	var cursor uint16
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
		e := make([]byte, eLen)
		if err := binary.Read(buf, binary.BigEndian, &e); err != nil {
			log.Println(err) //TODO::
		}
		eValue := bytes.NewBuffer(e)
		switch eType {
		case IETypeMeasurementMethod:
			createURR.MeasurementMethod = *DecodeMeasurementMethod(eValue, eLen)
		case IETypeMeasurementPeriod:
			createURR.MeasurementPeriod = *DecodeMeasurementPeriod(eValue, eLen)
		case IETypeReportingTriggers:
			createURR.ReportingTriggers = *DecodeReportingTriggers(eValue, eLen)
		case IETypeURRID:
			createURR.URRID = *DecodeURRID(eValue, eLen)
		default:
			log.Println("create urr err: unknown tlv type", eType) //TODO::
		}
		cursor = cursor + eLen + 4
	}
	return &createURR
}

func EncodeCreateURR(createURR CreateURR) []byte {
	ret := SetValue(createURR.EType, createURR.ELength)
	switch {
	case HasReportingTriggers(createURR.ReportingTriggers): //M
		SetValue(ret, EncodeReportingTriggers(createURR.ReportingTriggers))
		fallthrough
	case HasMeasurementMethod(createURR.MeasurementMethod): //M
		SetValue(ret, EncodeMeasurementMethod(createURR.MeasurementMethod))
		fallthrough
	case HasURRID(createURR.URRID): //M
		SetValue(ret, EncodeURRID(createURR.URRID))
		fallthrough
	case HasMeasurementPeriod(createURR.MeasurementPeriod):
		SetValue(ret, EncodeMeasurementPeriod(createURR.MeasurementPeriod))
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
