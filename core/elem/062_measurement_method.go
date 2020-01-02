package elem

type MeasurementMethod struct {
	EType   IEType
	ELength uint16
	DURAT   bool
	VOLUM   bool
	EVENT   bool
}

func DecodeMeasurementMethod(data []byte, len uint16) *MeasurementMethod {
	m := MeasurementMethod{
		EType:   IETypeMeasurementMethod,
		ELength: len,
	}
	flag := getValue(data, 1)[0]
	switch {
	case flag&0b00000001 == 1:
		m.DURAT = true
		fallthrough
	case flag&0b00000010>>1 == 1:
		m.VOLUM = true
		fallthrough
	case flag&0b00000100>>2 == 1:
		m.EVENT = true
	}
	return &m
}

func EncodeMeasurementMethod(m MeasurementMethod) []byte {
	var flag byte
	switch {
	case m.DURAT:
		flag |= 0b00000001
		fallthrough
	case m.VOLUM:
		flag |= 0b00000010
		fallthrough
	case m.EVENT:
		flag |= 0b00000100
	}
	return setValue(m.EType, m.ELength, flag)
}

//判断是否含有MeasurementMethod
func HasMeasurementMethod(m MeasurementMethod) bool {
	if m.EType == 0 {
		return false
	}
	return true
}
