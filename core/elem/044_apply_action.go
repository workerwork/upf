package elem

import "bytes"

type ActionType byte

const (
	_              ActionType = iota //0
	ActionTypeDROP                   //1
	ActionTypeFORW                   //2
	ActionTypeBUFF                   //3
	ActionTypeNOCP                   //4
	ActionTypeDUPL                   //5
)

type ApplyAction struct {
	EType   IEType
	ELength uint16
	Action  ActionType
}

func DecodeApplyAction(buf *bytes.Buffer, len uint16) *ApplyAction {
	return &ApplyAction{
		EType:   IETypeApplyAction,
		ELength: len,
		Action:  ActionType(getValue(buf, 1)[0]),
	}
}

func EncodeApplyAction(a ApplyAction) []byte {
	return setValue(a.EType, a.ELength, a.Action)
}

//判断是否含有ApplyAction
func HasApplyAction(a ApplyAction) bool {
	if a.EType == 0 {
		return false
	}
	return true
}
