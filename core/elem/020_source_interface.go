package elem

import "bytes"

/**************************************************************************
              			Bits
Octets          8	7	6	5	4	3	2	1
1 to 2          Type = 20 (decimal)
3 to 4          Length = n
5               Spare	Interface value
6 to (n+4)      These octet(s) is/are present only if explicitly specified
****************************************************************************/

type SourceInterface struct {
	EType     IEType
	ELength   uint16
	Interface InterfaceType //4bits
}

func DecodeSourceInterface(buf *bytes.Buffer, len uint16) *SourceInterface {
	return &SourceInterface{
		EType:     IETypeSourceInterface,
		ELength:   len,
		Interface: InterfaceType(getValue(buf, len)[0]),
	}
}

func EncodeSourceInterface(s SourceInterface) []byte {
	return setValue(s.EType, s.ELength, s.Interface)
}

//判断是否含有SourceInterface
func HasSourceInterface(s SourceInterface) bool {
	if s.EType == 0 {
		return false
	}
	return true
}
