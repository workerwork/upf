package elem

type PDRID struct {
	EType   IEType
	ELength uint16
	RuleID  []byte
}

func NewPDRID(ruleID []byte) *PDRID {
	return &PDRID{
		EType:   IETypePDRID,
		ELength: uint16(2),
		RuleID:  ruleID,
	}
}

func DecodePDRID(data []byte, len uint16) *PDRID {
	return &PDRID{
		EType:   IETypePDRID,
		ELength: len,
		RuleID:  getValue(data, len),
	}
}

func EncodePDRID(pdrID PDRID) []byte {
	return setValue(pdrID.EType, pdrID.ELength, pdrID.RuleID)
}

//判断是否含有PDRID
func HasPDRID(pdrID PDRID) bool {
	if pdrID.EType == 0 {
		return false
	}
	return true
}
