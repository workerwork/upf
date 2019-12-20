package elem

type FSEIDFlag byte

const (
	FSEIDFlagIPv6  FSEIDFlag = 1 //bits:0000 0001
	FSEIDFlagIPv4  FSEIDFlag = 2 //bits:0000 0010
	FSEIDFlagIPv46 FSEIDFlag = 3 //bits:0000 0011
)

type FSEID struct {
	EType    IEType
	ELength  uint16
	Flag     FSEIDFlag
	SEID     []byte //8byte
	IPv4Addr []byte //4byte
	IPv6Addr []byte //16byte
}

func DecodeFSEID(data []byte, len uint16) *FSEID {
	f := FSEID{
		EType:   IETypeFSEID,
		ELength: len,
		Flag:    FSEIDFlag(getValue(data, 1)[0]),
		SEID:    getValue(data, 8),
	}
	switch f.Flag {
	case FSEIDFlagIPv4:
		f.IPv4Addr = getValue(data, 4)
	case FSEIDFlagIPv6:
		f.IPv6Addr = getValue(data, 16)
	case FSEIDFlagIPv46:
		f.IPv4Addr = getValue(data, 4)
		f.IPv6Addr = getValue(data, 16)
	}
	return &f
}

func EncodeFSEID(f FSEID) []byte {
	return setValue(f.EType, f.ELength, f.Flag, f.SEID, f.IPv4Addr, f.IPv6Addr)
}

//判断是否含有FSEID
func HasFSEID(f FSEID) bool {
	if f.EType == 0 {
		return false
	}
	return true
}
