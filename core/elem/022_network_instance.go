package elem

import "bytes"

type NetworkInstance struct {
	EType           IEType
	ELength         uint16
	NetworkInstance []byte
}

func DecodeNetworkInstance(buf *bytes.Buffer, len uint16) *NetworkInstance {
	return &NetworkInstance{
		EType:           IETypeNetworkInstance,
		ELength:         len,
		NetworkInstance: getValue(buf, len),
	}
}

func EncodeNetworkInstance(n NetworkInstance) []byte {
	return setValue(n.EType, n.ELength, n.NetworkInstance)
}

//判断是否含有NetworkInstance
func HasNetworkInstance(n NetworkInstance) bool {
	if n.EType == 0 {
		return false
	}
	return true
}
