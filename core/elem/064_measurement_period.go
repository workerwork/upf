package elem

import "bytes"

type MeasurementPeriod struct {
	EType             IEType
	ELength           uint16
	MeasurementPeriod []byte //4byte
}

func DecodeMeasurementPeriod(buf *bytes.Buffer, len uint16) *MeasurementPeriod {
	return &MeasurementPeriod{
		EType:             IETypeMeasurementPeriod,
		ELength:           len,
		MeasurementPeriod: GetValue(buf, 4),
	}
}

func EncodeMeasurementPeriod(m MeasurementPeriod) []byte {
	return SetValue(m.EType, m.ELength, m.MeasurementPeriod)
}

//判断是否含有MeasurementPeriod
func HasMeasurementPeriod(m MeasurementPeriod) bool {
	if m.EType == 0 {
		return false
	}
	return true
}
