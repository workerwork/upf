package elem

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

func DecodeOuterHeaderCreation(data []byte, len uint16) *OuterHeaderCreation {
	return &OuterHeaderCreation{
		EType:                          IETypeOuterHeaderCreation,
		ELength:                        len,
		OuterHeaderCreationDescription: getValue(data, 2),
		TEID:                           getValue(data, 4),
		IPv4Addr:                       getValue(data, 4),
		IPv6Addr:                       getValue(data, 16),
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
