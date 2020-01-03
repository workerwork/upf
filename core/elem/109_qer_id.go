package elem

import "bytes"

type QERID struct {
	EType   IEType
	ELength uint16
	QERID   []byte //4byte
}

func DecodeQERID(buf *bytes.Buffer, len uint16) *QERID {
	return &QERID{
		EType:   IETypeQERID,
		ELength: len,
		QERID:   getValue(buf, 4),
	}
}

func EncodeQERID(q QERID) []byte {
	return setValue(q.EType, q.ELength, q.QERID)
}

//判断是否含有QERID
func HasQERID(q QERID) bool {
	if q.EType == 0 {
		return false
	}
	return true
}
