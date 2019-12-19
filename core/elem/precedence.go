package elem

type Precedence struct {
	EType      IEType
	ELength    uint16
	Precedence []byte
}

func NewPrecedence(p []byte) *Precedence {
	return &Precedence{
		EType:      IETypePrecedence,
		ELength:    uint16(4),
		Precedence: p,
	}
}

func DecodePrecedence(data []byte, len uint16) *Precedence {
	return &Precedence{
		EType:      IETypePrecedence,
		ELength:    len,
		Precedence: getValue(data, len),
	}
}

func EncodePrecedence(p Precedence) []byte {
	return setBuffer(p.EType, p.ELength, p.Precedence)
}

//判断是否含有Precedence
func HasPrecedence(p Precedence) bool {
	if p.EType == 0 {
		return false
	}
	return true
}
