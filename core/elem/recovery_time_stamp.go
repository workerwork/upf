package elem

type RecoveryTimeStamp struct {
	EType             IEType
	ELength           uint16
	RecoveryTimeStamp []byte
}

func NewRecoveryTimeStamp(rts []byte) *RecoveryTimeStamp {
	return &RecoveryTimeStamp{
		EType:             IETYPE_RECOVERY_TIME_STAMP,
		ELength:           uint16(4),
		RecoveryTimeStamp: rts,
	}
}

func DecodeRecoveryTimeStamp(data []byte, len uint16) *RecoveryTimeStamp {
	return &RecoveryTimeStamp{
		EType:             IETYPE_RECOVERY_TIME_STAMP,
		ELength:           len,
		RecoveryTimeStamp: getValue(data, len),
	}
}

func EncodeRecoveryTimeStamp(rts RecoveryTimeStamp) []byte {
	return setBuffer(rts.EType, rts.ELength, rts.RecoveryTimeStamp)
}

//判断是否含有RecoveryTimeStamp
func HasRecoveryTimeStamp(rts RecoveryTimeStamp) bool {
	if rts.EType == 0 {
		return false
	}
	return true
}
