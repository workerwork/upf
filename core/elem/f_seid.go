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
	SEID     []byte	//8byte
	IPv4Addr []byte	//4byte
	IPv6Addr []byte //16byte
}

func NewIPv4FSEID(seid []byte, ipv4 uint16) *FSEID {
	return &FSEID{
		EType:    IETypeFSEID,
		ELength:  uint16(14),
		Flag:     FSEIDFlagIPv4,
		SEID:     seid,
		IPv4Addr: ipv4,
	}
}

func DecodeFSEID(data []byte, len uint16) *FSEID {
	var fseid FSEID
	fseid.EType = IETypeFSEID
	fseid.ELength = len
	fseid.Flag = FSEIDFlag(getValue(data, 1)[0])
	fseid.SEID = getValue(data, 8)
	switch fseid.Flag {
	case FSEIDFlagIPv4:
		fseid.IPv4Addr = getValue(data, 4)
	case FSEIDFlagIPv6:
		fseid.IPv6Addr = getValue(data, 16)
	case FSEIDFlagIPv46:
		fseid.IPv4Addr = getValue(data, 4)
		fseid.IPv6Addr = getValue(data, 16)
	}
	return &fseid
}

func EncodeFSEID(fseid FSEID) []byte {
	return setValue(fseid.EType, fseid.ELength, fseid.Flag, fseid.SEID, fseid.IPv4Addr, fseid.IPv6Addr)
}

//判断是否含有FSEID
func HasFSEID(fseid FSEID) bool {
	if fseid.EType == 0 {
		return false
	}
	return true
}
