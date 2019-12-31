package elem

type ActionType byte

const (
	_    ActionType = iota //0
	DROP                   //1
	FORW                   //2
	BUFF                   //3
	NOCP                   //4
	DUPL                   //5
)

type ApplyAction struct {
	EType   IEType
	ELength uint16
	Action  ActionType
}

func DecodeApplyAction(data []byte, len uint16) *ApplyAction {
	return &ApplyAction{
		EType:   IETypeApplyAction,
		ELength: len,
		Action:  ActionType(getValue(data, 1)[0]),
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
