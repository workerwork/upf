package elem

import "bytes"

//type ReportingTriggersFlag byte

type ReportingTriggers struct {
	EType   IEType
	ELength uint16
	PERIO   bool
	VOLTH   bool
	TIMTH   bool
	QUHTI   bool
	START   bool
	STOPT   bool
	DROTH   bool
	LIUSA   bool
	VOLQU   bool
	TIMQU   bool
	ENVCL   bool
	MACAR   bool
	EVETH   bool
	EVEQU   bool
}

func DecodeReportingTriggers(buf *bytes.Buffer, len uint16) *ReportingTriggers {
	r := ReportingTriggers{
		EType:   IETypeReportingTriggers,
		ELength: len,
	}
	flag1 := GetValue(buf, 1)[0]
	switch {
	case flag1&0b00000001 != 0:
		r.PERIO = true
		fallthrough
	case flag1&0b00000010 != 0:
		r.VOLTH = true
		fallthrough
	case flag1&0b00000100 != 0:
		r.TIMTH = true
		fallthrough
	case flag1&0b00001000 != 0:
		r.QUHTI = true
		fallthrough
	case flag1&0b00010000 != 0:
		r.START = true
		fallthrough
	case flag1&0b00100000 != 0:
		r.STOPT = true
		fallthrough
	case flag1&0b01000000 != 0:
		r.DROTH = true
		fallthrough
	case flag1&0b10000000 != 0:
		r.LIUSA = true
	}
	flag2 := GetValue(buf, 1)[0]
	switch {
	case flag2&0b00000001 != 0:
		r.VOLQU = true
		fallthrough
	case flag2&0b00000010 != 0:
		r.TIMQU = true
		fallthrough
	case flag2&0b00000100 != 0:
		r.ENVCL = true
		fallthrough
	case flag2&0b00001000 != 0:
		r.MACAR = true
		fallthrough
	case flag2&0b00010000 != 0:
		r.EVETH = true
		fallthrough
	case flag2&0b00100000 != 0:
		r.EVEQU = true
	}
	return &r
}

func EncodeReportingTriggers(r ReportingTriggers) []byte {
	var flag1, flag2 byte
	switch {
	case r.PERIO:
		flag1 |= 0b00000001
		fallthrough
	case r.VOLTH:
		flag1 |= 0b00000010
		fallthrough
	case r.TIMTH:
		flag1 |= 0b00000100
		fallthrough
	case r.QUHTI:
		flag1 |= 0b00001000
		fallthrough
	case r.START:
		flag1 |= 0b00010000
		fallthrough
	case r.STOPT:
		flag1 |= 0b00100000
		fallthrough
	case r.DROTH:
		flag1 |= 0b01000000
		fallthrough
	case r.LIUSA:
		flag1 |= 0b10000000
	}
	switch {
	case r.VOLQU:
		flag2 |= 0b00000001
		fallthrough
	case r.TIMQU:
		flag2 |= 0b00000010
		fallthrough
	case r.ENVCL:
		flag2 |= 0b00000100
		fallthrough
	case r.MACAR:
		flag2 |= 0b00001000
		fallthrough
	case r.EVETH:
		flag2 |= 0b00010000
		fallthrough
	case r.EVEQU:
		flag2 |= 0b00100000
	}
	return SetValue(r.EType, r.ELength, flag1, flag2)
}

//判断是否含有ReportingTriggers
func HasReportingTriggers(r ReportingTriggers) bool {
	if r.EType == 0 {
		return false
	}
	return true
}
