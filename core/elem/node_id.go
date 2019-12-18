package elem

type NodeIDType byte

const (
	IPv4Addr NodeIDType = 0
	IPv6Addr NodeIDType = 1
	FQDN     NodeIDType = 2
)

type NodeID struct {
	EType   IEType
	ELength uint16
	NodeIDType
	NodeID []byte
}

func NewIPv4NodeID(ipv4 []byte) *NodeID {
	return &NodeID{
		EType:   IETYPE_NODE_ID,
		ELength: uint16(4),
		NodeID:  ipv4,
	}
}

func DecodeNodeID(data []byte, len uint16) *NodeID {
	return &NodeID{
		EType:      IETYPE_NODE_ID,
		ELength:    len,
		NodeIDType: NodeIDType(getValue(data, 1)[0]),
		NodeID:     getValue(data, len-1),
	}
}

func EncodeNodeID(nodeID NodeID) []byte {
	return setBuffer(nodeID.EType, nodeID.ELength, setValue(byte(nodeID.NodeIDType), nodeID.NodeID))
}

//判断是否含有NodeID
func HasNodeID(nodeID NodeID) bool {
	if nodeID.EType == 0 {
		return false
	}
	return true
}
