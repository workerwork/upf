package elem

type MBR struct {
	EType   IEType
	ELength uint16
	ULMBR   []byte //5bytes
	DLMBR   []byte //5bytes
}

func DecodeMBR(data []byte, len uint16) *MBR {
	return &MBR{
		EType:   IETypeMBR,
		ELength: len,
		ULMBR:   getValue(data, 5),
		DLMBR:   getValue(data, 5),
	}
}

func EncodeMBR(m MBR) []byte {
	return setValue(g.EType, g.ELength, m.ULMBR, m.DLMBR)
}

//判断是否含有MBR
func HasMBR(m MBR) bool {
	if m.EType == 0 {
		return false
	}
	return true
}
