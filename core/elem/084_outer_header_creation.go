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
		OuterHeaderCreationDescription: getValue(buf, 2),
	}
	flag := o.OuterHeaderCreationDescription[1]
	if flag&0b00000001 == 1 || flag&0b00000010>>1 == 1 {
		o.TEID = getValue(buf, 4)
	}
	if flag&0b00000001 == 1 || flag&0b00000100>>2 == 1 || flag&0b00010000>>4 == 1 {
		o.IPv4Addr = getValue(buf, 4)
	}
	if flag&0b00000010>>1 == 1 || flag&0b00001000>>3 == 1 || flag&0b00100000>>5 == 1 {
		o.IPv6Addr = getValue(buf, 16)
	}
	if flag&0b01000000>>6 == 1 {
		o.CTAG = getValue(buf, 3)
	}
	if flag&0b10000000>>7 == 1 {
		o.STAG = getValue(buf, 3)
	}
	return &o
}

func EncodeOuterHeaderCreation(o OuterHeaderCreation) *bytes.Buffer {
	return SetValue(o.EType, o.ELength, o.OuterHeaderCreationDescription, o.TEID, o.IPv4Addr, o.IPv6Addr, o.CTAG, o.STAG)
}

//判断是否含有OuterHeaderCreation
func HasOuterHeaderCreation(o OuterHeaderCreation) bool {
	if o.EType == 0 {
		return false
	}
	return true
}
