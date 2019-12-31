package elem

type MeasurementPeriod struct {
	EType   IEType
	ELength uint16
	RuleID  []byte //2byte
}

func DecodeMeasurementPeriod(data []byte, len uint16) *MeasurementPeriod {
	return &MeasurementPeriod{
		EType:   IETypePDRID,
		ELength: len,
		RuleID:  getValue(data, 2),
	}
}

func EncodeMeasurementPeriod(m MeasurementPeriod) []byte {
	return setValue(m.EType, m.ELength, m.RuleID)
}

//判断是否含有MeasurementPeriod
func HasMeasurementPeriod(m MeasurementPeriod) bool {
	if m.EType == 0 {
		return false
	}
	return true
}

