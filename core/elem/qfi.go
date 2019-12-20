package elem


type QFI struct {
	EType     IEType
	ELength   uint16
	QFI []byte
}

func DecodeQFI(data []byte, len uint16) *QFI {
	return &QFI{
		EType: IETypeQFI,
		ELength: len,
		QFI: getValue(data,len),
	}
}

func EncodeQFI(qfi QFI) []byte {
	return setValue(qfi.EType, qfi.ELength, qfi.QFI)
}

//判断是否含有SourceInterface
func HasQFI(qfi QFI) bool {
	if qfi.EType == 0 {
		return false
	}
	return true
}
