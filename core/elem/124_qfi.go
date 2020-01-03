package elem

import "bytes"

type QFI struct {
	EType   IEType
	ELength uint16
	QFI     byte
}

func DecodeQFI(buf *bytes.Buffer, len uint16) *QFI {
	return &QFI{
		EType:   IETypeQFI,
		ELength: len,
		QFI:     getValue(buf, len)[0],
	}
}

func EncodeQFI(q QFI) []byte {
	return setValue(q.EType, q.ELength, q.QFI)
}

//判断是否含有SourceInterface
func HasQFI(q QFI) bool {
	if q.EType == 0 {
		return false
	}
	return true
}
