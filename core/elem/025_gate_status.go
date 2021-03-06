package elem

import "bytes"

type GateStatusType byte

const (
	GateStatusTypeOPEN   GateStatusType = iota //0
	GateStatusTypeCLOSED                       //1
)

type GateStatus struct {
	EType   IEType
	ELength uint16
	DLGate  GateStatusType
	ULGate  GateStatusType
}

func DecodeGateStatus(buf *bytes.Buffer, len uint16) *GateStatus {
	g := GateStatus{
		EType:   IETypeGateStatus,
		ELength: len,
		DLGate:  GateStatusTypeCLOSED,
		ULGate:  GateStatusTypeCLOSED,
	}
	flag := GetValue(buf, 1)[0]
	if flag&0b00000001 == 0 {
		g.DLGate = GateStatusTypeOPEN
	}
	if flag&0b00000100>>2 == 0 {
		g.ULGate = GateStatusTypeOPEN
	}
	return &g
}

func EncodeGateStatus(g GateStatus) []byte {
	var flag byte
	if g.DLGate == GateStatusTypeCLOSED {
		flag |= 0b00000001
	}
	if g.ULGate == GateStatusTypeCLOSED {
		flag |= 0b00000100
	}
	return SetValue(g.EType, g.ELength, flag)
}

//判断是否含有GateStatus
func HasGateStatus(g GateStatus) bool {
	if g.EType == 0 {
		return false
	}
	return true
}
