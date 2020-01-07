package elem

import "bytes"

type PDRID struct {
	EType   IEType
	ELength uint16
	RuleID  []byte //2byte
}

func NewPDRID(r []byte) *PDRID {
	return &PDRID{
		EType:   IETypePDRID,
		ELength: uint16(2),
		RuleID:  r,
	}
}

func DecodePDRID(buf *bytes.Buffer, len uint16) *PDRID {
	return &PDRID{
		EType:   IETypePDRID,
		ELength: len,
		RuleID:  GetValue(buf, 2),
	}
}

func EncodePDRID(p PDRID) []byte {
	return SetValue(p.EType, p.ELength, p.RuleID)
}

//判断是否含有PDRID
func HasPDRID(p PDRID) bool {
	if p.EType == 0 {
		return false
	}
	return true
}
