package elem

type MeasurementPeriod struct {
	EType             IEType
	ELength           uint16
	MeasurementPeriod []byte //4byte
}

func DecodeMeasurementPeriod(data []byte, len uint16) *MeasurementPeriod {
	return &MeasurementPeriod{
		EType:             IETypeMeasurementPeriod,
		ELength:           len,
		MeasurementPeriod: getValue(data, 4),
	}
}

func EncodeMeasurementPeriod(m MeasurementPeriod) []byte {
	return setValue(m.EType, m.ELength, m.MeasurementPeriod)
}

//判断是否含有MeasurementPeriod
func HasMeasurementPeriod(m MeasurementPeriod) bool {
	if m.EType == 0 {
		return false
	}
	return true
}
