package elem

import "bytes"

type PDNTypeType byte

const (
	_                   PDNTypeType = iota //0
	PDNTypeTypeIPv4                        //1
	PDNTypeTypeIPv6                        //2
	PDNTypeTypeIPv4v6                      //3
	PDNTypeTypeNonIP                       //4
	PDNTypeTypeEthernet                    //5
)

type PDNType struct {
	EType   IEType
	ELength uint16
	PDNType PDNTypeType
}

func DecodePDNType(buf *bytes.Buffer, len uint16) *PDNType {
	return &PDNType{
		EType:   IETypePDNType,
		ELength: len,
		PDNType: PDNTypeType(getValue(buf, 1)[0]),
	}
}

func EncodePDNType(p PDNType) []byte {
	return setValue(p.EType, p.ELength, p.PDNType)
}

//判断是否含有PDNType
func HasPDNType(p PDNType) bool {
	if p.EType == 0 {
		return false
	}
	return true
}
