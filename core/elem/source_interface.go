package elem

/**************************************************************************
              			Bits
Octets		8	7	6	5	4	3	2	1
1 to 2			Type = 20 (decimal)
3 to 4			Length = n
5				Spare	Interface value
6 to (n+4)		These octet(s) is/are present only if explicitly specified
****************************************************************************/
type InterfaceType byte

const (
	InterfaceTypeAccess       InterfaceType = iota //0
	InterfaceTypeCore                              //1
	InterfaceTypeSGiLAN                            //2
	InterfaceTypeCPFunction                        //3
	InterfaceType5GVNInternal                      //4
)

type SourceInterface struct {
	EType     IEType
	ELength   uint16
	Interface InterfaceType
}

func DecodeSourceInterface(data []byte, len uint16) *SourceInterface {
	return &SourceInterface{
		EType: IETypeFTEID,
		ELength: len,
		Interface: InterfaceType(getValue(data,len)[0]),
	}
}

func EncodeSourceInterface(s SourceInterface) []byte {
	return setValue(s.EType, s.ELength, s.Interface)
}

//判断是否含有SourceInterface
func HasSourceInterface(s SourceInterface) bool {
	if s.EType == 0 {
		return false
	}
	return true
}