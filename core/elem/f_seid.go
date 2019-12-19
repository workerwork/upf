package elem

type IPType byte

const (
	IPTypeIPv6  IPType = 1
	IPTypeIPv4  IPType = 2
	IPTypeIPv46 IPType = 3
)

type FSEID struct {
	EType    IEType
	ELength  uint16
	IPType   IPType
	SEID     []byte
	IPv4Addr []byte
	IPv6Addr []byte
}

func NewIPv4FSEID(seid []byte, ipv4 []byte) *FSEID {
	return &FSEID{
		EType:    IETypeFSEID,
		ELength:  uint16(13),
		IPType:   IPTypeIPv4,
		SEID:     seid,
		IPv4Addr: ipv4,
	}
}

func DecodeFSEID(data []byte, len uint16) *FSEID {
	var fseid FSEID
	fseid.EType = IETypeFSEID
	fseid.ELength = len
	fseid.IPType = IPType(getValue(data, 1)[0])
	fseid.SEID = getValue(data, 8)
	switch fseid.IPType {
	case IPTypeIPv4:
		fseid.IPv4Addr = getValue(data, 4)
	case IPTypeIPv6:
		fseid.IPv6Addr = getValue(data, 16)
	case IPTypeIPv46:
		fseid.IPv4Addr = getValue(data, 4)
		fseid.IPv6Addr = getValue(data, 16)
	}
	return &fseid
}

func EncodeFSEID(fseid FSEID) []byte {
	return setValue(fseid.EType, fseid.ELength, fseid.IPType, fseid.SEID, fseid.IPv4Addr, fseid.IPv6Addr)
}

//判断是否含有FSEID
func HasFSEID(fseid FSEID) bool {
	if fseid.EType == 0 {
		return false
	}
	return true
}
