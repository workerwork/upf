package elem

import "bytes"

type OuterHeaderRemoval struct {
	EType                         IEType
	ELength                       uint16
	OuterHeaderRemovalDescription byte
	GTPUExtensionHeaderDeletion   byte
}

func DecodeOuterHeaderRemoval(buf *bytes.Buffer, len uint16) *OuterHeaderRemoval {
	o := OuterHeaderRemoval{
		EType:                         IETypeOuterHeaderRemoval,
		ELength:                       len,
		OuterHeaderRemovalDescription: getValue(buf, 1)[0],
	}
	if len > 1 {
		o.GTPUExtensionHeaderDeletion = getValue(buf, 1)[0]
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
