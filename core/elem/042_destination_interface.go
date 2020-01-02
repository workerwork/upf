package elem

type DestinationInterface struct {
	EType     IEType
	ELength   uint16
	Interface InterfaceType //4bits
}

func DecodeDestinationInterface(data []byte, len uint16) *DestinationInterface {
	return &DestinationInterface{
		EType:     IETypeDestinationInterface,
		ELength:   len,
		Interface: InterfaceType(getValue(data, len)[0]),
	}
}

func EncodeDestinationInterface(d DestinationInterface) []byte {
	return setValue(d.EType, d.ELength, d.Interface)
}

//判断是否含有DestinationInterface
func HasDestinationInterface(d DestinationInterface) bool {
	if d.EType == 0 {
		return false
	}
	return true
}