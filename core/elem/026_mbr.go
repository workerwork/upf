package elem

import "bytes"

type MBR struct {
	EType   IEType
	ELength uint16
	ULMBR   []byte //5bytes
	DLMBR   []byte //5bytes
}

func DecodeMBR(buf *bytes.Buffer, len uint16) *MBR {
	return &MBR{
		EType:   IETypeMBR,
		ELength: len,
		ULMBR:   getValue(buf, 5),
		DLMBR:   getValue(buf, 5),
	}
}

func EncodeMBR(m MBR) []byte {
	return SetValue(m.EType, m.ELength, m.ULMBR, m.DLMBR)
}

//判断是否含有MBR
func HasMBR(m MBR) bool {
	if m.EType == 0 {
		return false
	}
	return true
}
