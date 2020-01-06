package elem

import "bytes"

type NodeIDType byte

const (
	NodeIDTypeIPv4Addr NodeIDType = 0
	NodeIDTypeIPv6Addr NodeIDType = 1
	NodeIDTypeFQDN     NodeIDType = 2
)

type NodeID struct {
	EType   IEType
	ELength uint16
	NodeIDType
	NodeID []byte
}

func NewIPv4NodeID(ipv4 []byte) *NodeID {
	return &NodeID{
		EType:      IETypeNodeID,
		ELength:    uint16(4),
		NodeIDType: NodeIDTypeIPv4Addr,
		NodeID:     ipv4,
	}
}

func DecodeNodeID(buf *bytes.Buffer, len uint16) *NodeID {
	return &NodeID{
		EType:      IETypeNodeID,
		ELength:    len,
		NodeIDType: NodeIDType(getValue(buf, 1)[0]),
		NodeID:     getValue(buf, len-1),
	}
}

func EncodeNodeID(n NodeID) *bytes.Buffer {
	return SetValue(n.EType, n.ELength, n.NodeIDType, n.NodeID)
}

//判断是否含有NodeID
func HasNodeID(n NodeID) bool {
	if n.EType == 0 {
		return false
	}
	return true
}
