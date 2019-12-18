package msg

func getDataLen(msg *Msg) uint16 {
	if msg.S {
		return msg.Length - (16 - 4)
	}
	return msg.Length - (8 - 4)
}
