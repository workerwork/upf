package elem

import (
	"bytes"
)

type FTEIDFlag byte

type FTEID struct {
	EType    IEType
	ELength  uint16
	Flag     FTEIDFlag
	TEID     []byte //4byte
	IPv4Addr []byte //4byte
	IPv6Addr []byte //16byte
	ChooseID byte
}

func DecodeFTEID(buf *bytes.Buffer, len uint16) *FTEID {
	f := FTEID{
		EType:   IETypeFTEID,
		ELength: len,
		Flag:    FTEIDFlag(GetValue(buf, 1)[0]),
	}
	if f.Flag&0b00000100 != 0 { //CH=1
		f.ChooseID = GetValue(buf, 1)[0]
	} else {
		f.TEID = GetValue(buf, 4)
		switch {
		case f.Flag&0b00000001 != 0: //v4=1
			f.IPv4Addr = GetValue(buf, 4)
			fallthrough
		case f.Flag&0b00000010 != 0: //v6=1
			f.IPv6Addr = GetValue(buf, 16)
			fallthrough
		case f.Flag&0b00001000 != 0: //CHID=1
			f.ChooseID = GetValue(buf, 1)[0]
		}
	}
	return &f
}

func EncodeFTEID(f FTEID) []byte {
	ret := SetValue(f.EType, f.ELength, f.Flag, f.TEID, f.IPv4Addr, f.IPv6Addr)
	if f.Flag&0b00000100>>2 == 0 && f.Flag&0b00001000 != 0 {
		SetValue(ret, f.ChooseID)
	}
	return ret
}

//判断是否含有FTEID
func HasFTEID(f FTEID) bool {
	if f.EType == 0 {
		return false
	}
	return true
}
