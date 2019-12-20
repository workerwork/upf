package elem

type URRID struct {
	EType   IEType
	ELength uint16
	URRID   []byte //4byte //TODO::
}

func DecodeURRID(data []byte, len uint16) *URRID {
	return &URRID{
		EType:   IETypeURRID,
		ELength: len,
		URRID:   getValue(data, len),
	}
}

func EncodeURRID(u URRID) []byte {
	return setValue(u.EType, u.ELength, u.URRID)
}

//判断是否含有URRID
func HasURRID(u URRID) bool {
	if u.EType == 0 {
		return false
	}
	return true
}
