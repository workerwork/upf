package elem

import "bytes"

type URRID struct {
	EType   IEType
	ELength uint16
	URRID   []byte //4byte //TODO::
}

func DecodeURRID(buf *bytes.Buffer, len uint16) *URRID {
	return &URRID{
		EType:   IETypeURRID,
		ELength: len,
		URRID:   getValue(buf, 4),
	}
}

func EncodeURRID(u URRID) *bytes.Buffer {
	return SetValue(u.EType, u.ELength, u.URRID)
}

//判断是否含有URRID
func HasURRID(u URRID) bool {
	if u.EType == 0 {
		return false
	}
	return true
}
