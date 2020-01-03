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
	return &OuterHeaderCreation{
		EType:                          IETypeOuterHeaderCreation,
		ELength:                        len,
		OuterHeaderCreationDescription: getValue(buf, 2),
		TEID:                           getValue(buf, 4),
		IPv4Addr:                       getValue(buf, 4),
		IPv6Addr:                       getValue(buf, 16),
		//TODO::
	}
}

func EncodeOuterHeaderCreation(pdrID OuterHeaderCreation) []byte {
	return []byte{}
}

//判断是否含有OuterHeaderCreation
func HasOuterHeaderCreation(o OuterHeaderCreation) bool {
	if o.EType == 0 {
		return false
	}
	return true
}
