package elem

import "bytes"

type Precedence struct {
	EType      IEType
	ELength    uint16
	Precedence []byte //4byte
}

func NewPrecedence(p []byte) *Precedence {
	return &Precedence{
		EType:      IETypePrecedence,
		ELength:    uint16(4),
		Precedence: p,
	}
}

func DecodePrecedence(buf *bytes.Buffer, len uint16) *Precedence {
	return &Precedence{
		EType:      IETypePrecedence,
		ELength:    len,
		Precedence: getValue(buf, len),
	}
}

func EncodePrecedence(p Precedence) []byte {
	return SetValue(p.EType, p.ELength, p.Precedence)
}

//判断是否含有Precedence
func HasPrecedence(p Precedence) bool {
	if p.EType == 0 {
		return false
	}
	return true
}
