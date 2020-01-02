package elem

type RecoveryTimeStamp struct {
	EType             IEType
	ELength           uint16
	RecoveryTimeStamp []byte //4byte
}

func NewRecoveryTimeStamp(r []byte) *RecoveryTimeStamp {
	return &RecoveryTimeStamp{
		EType:             IETypeRecoveryTimeStamp,
		ELength:           uint16(4),
		RecoveryTimeStamp: r,
	}
}

func DecodeRecoveryTimeStamp(data []byte, len uint16) *RecoveryTimeStamp {
	return &RecoveryTimeStamp{
		EType:             IETypeRecoveryTimeStamp,
		ELength:           len,
		RecoveryTimeStamp: getValue(data, len),
	}
}

func EncodeRecoveryTimeStamp(r RecoveryTimeStamp) []byte {
	return setValue(r.EType, r.ELength, r.RecoveryTimeStamp)
}

//判断是否含有RecoveryTimeStamp
func HasRecoveryTimeStamp(r RecoveryTimeStamp) bool {
	if r.EType == 0 {
		return false
	}
	return true
}
