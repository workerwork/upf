package elem

type OuterHeaderRemoval struct {
	EType                         IEType
	ELength                       uint16
	OuterHeaderRemovalDescription []byte //2byte
	GTPUExtensionHeaderDeletion   []byte //4byte

}

func DecodeOuterHeaderRemoval(data []byte, len uint16) *OuterHeaderRemoval {
	o := OuterHeaderRemoval{
		EType:                         IETypeOuterHeaderCreation,
		ELength:                       len,
		OuterHeaderRemovalDescription: getValue(data, 1),
	}
	if len > 1 {
		o.GTPUExtensionHeaderDeletion = getValue(data, 1)
	}
	return &o
}

func EncodeOuterHeaderRemoval(o OuterHeaderRemoval) []byte {
	return setValue(o.EType, o.ELength, o.OuterHeaderRemovalDescription, o.GTPUExtensionHeaderDeletion)
}

//判断是否含有OuterHeaderRemoval
func HasOuterHeaderRemoval(o OuterHeaderRemoval) bool {
	if o.EType == 0 {
		return false
	}
	return true
}
