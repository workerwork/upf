package elem

import "bytes"

type FARID struct {
	EType   IEType
	ELength uint16
	FARID   []byte //4byte
}

func DecodeFARID(buf *bytes.Buffer, len uint16) *FARID {
	return &FARID{
		EType:   IETypeFARID,
		ELength: len,
		FARID:   getValue(buf, 4),
	}
}

func EncodeFARID(f FARID) *bytes.Buffer {
	return SetValue(f.EType, f.ELength, f.FARID)
}

//判断是否含有FARID
func HasFARID(f FARID) bool {
	if f.EType == 0 {
		return false
	}
	return true
}
