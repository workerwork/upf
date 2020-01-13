package elem

import "bytes"

type OuterHeaderCreation struct {
	EType                          IEType
	ELength                        uint16
	OuterHeaderCreationDescription []byte //2byte
	TEID                           []byte //4byte
	IPv4Addr                       []byte //4byte
	IPv6Addr                       []byte //16byte
	PortNumber                     []byte //2byte
	CTAG                           []byte //3byte
	STAG                           []byte //3byte
}

func DecodeOuterHeaderCreation(buf *bytes.Buffer, len uint16) *OuterHeaderCreation {
	o := OuterHeaderCreation{
		EType:                          IETypeOuterHeaderCreation,
		ELength:                        len,
		OuterHeaderCreationDescription: GetValue(buf, 2),
	}
	flag := o.OuterHeaderCreationDescription[1]
	switch {
	case flag&0b00000011 != 0:
		o.TEID = GetValue(buf, 4)
		fallthrough
	case flag&0b00010101 != 0:
		o.IPv4Addr = GetValue(buf, 4)
		fallthrough
	case flag&0b00101010 != 0:
		o.IPv6Addr = GetValue(buf, 16)
		fallthrough
	case flag&0b00001111 != 0:
		o.PortNumber = GetValue(buf, 1)
		fallthrough
	case flag&0b01000000 != 0:
		o.CTAG = GetValue(buf, 3)
		fallthrough
	case flag&0b10000000 != 0:
		o.STAG = GetValue(buf, 3)
	}
	return &o
}

func EncodeOuterHeaderCreation(o OuterHeaderCreation) []byte {
	return SetValue(o.EType, o.ELength, o.OuterHeaderCreationDescription, o.TEID, o.IPv4Addr, o.IPv6Addr, o.PortNumber, o.CTAG, o.STAG)
}

//判断是否含有OuterHeaderCreation
func HasOuterHeaderCreation(o OuterHeaderCreation) bool {
	if o.EType == 0 {
		return false
	}
	return true
}
