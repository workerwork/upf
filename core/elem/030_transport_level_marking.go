package elem

import "bytes"

type TransportLevelMarking struct {
	EType             IEType
	ELength           uint16
	TosOrTrafficClass []byte //2byte
}

func DecodeTransportLevelMarking(buf *bytes.Buffer, len uint16) *TransportLevelMarking {
	return &TransportLevelMarking{
		EType:             IETypeTransportLevelMarking,
		ELength:           len,
		TosOrTrafficClass: getValue(buf, 2),
	}
}

func EncodeTransportLevelMarking(t TransportLevelMarking) []byte {
	return setValue(t.EType, t.ELength, t.TosOrTrafficClass)
}

//判断是否含有TransportLevelMarking
func HasTransportLevelMarking(t TransportLevelMarking) bool {
	if t.EType == 0 {
		return false
	}
	return true
}
