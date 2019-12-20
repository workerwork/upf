package elem

type URRID struct {
	EType   IEType
	ELength uint16
	URRID  []byte	//4byte //TODO::
}


func DecodeURRID(data []byte, len uint16) *URRID {
	return &URRID{
		EType:   IETypeURRID,
		ELength: len,
		URRID:  getValue(data, len),
	}
}

func EncodeURRID(urrID URRID) []byte {
	return setValue(urrID.EType, urrID.ELength, urrID.URRID)
}

//判断是否含有URRID
func HasURRID(urrID URRID) bool {
	if urrID.EType == 0 {
		return false
	}
	return true
}


