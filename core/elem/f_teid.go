package elem

type FTEIDFlag byte

/**
const (
	FTEIDFlagIPv4    FTEIDFlag = 1 //bits:0000 0001
	FTEIDFlagIPv6    FTEIDFlag = 2 //bits:0000 0010
	FTEIDFlagIPv46   FTEIDFlag = 3 //bits:0000 0011
	FTEIDFlagCH      FTEIDFlag = 4 //bits:0000 0100
	FTEIDFlagCHIPv4  FTEIDFlag = 5 //bits:0000 0100
	FTEIDFlagCHIPv6  FTEIDFlag = 6 //bits:0000 0100
	FTEIDFlagCHIPv46 FTEIDFlag = 7 //bits:0000 0100
)*/

type FTEID struct {
	EType    IEType
	ELength  uint16
	Flag     FTEIDFlag
	TEID     []byte
	IPv4Addr []byte
	IPv6Addr []byte
	ChooseID []byte
}

func NewIPv4TEID(teid []byte, ipv4 []byte) *FTEID {
	return &FTEID{}
}

func DecodeTEID(data []byte, len uint16) *FTEID {
	var fteid FTEID
	fteid.EType = IETypeFTEID
	fteid.ELength = len
	fteid.Flag = FTEIDFlag(getValue(data, 1)[0])
	if fteid.Flag&0b00000100 == 1 { //CH=1
		fteid.ChooseID = getValue(data, 1)
	} else {
		fteid.TEID = getValue(data, 4)
		if fteid.Flag&0b00000001 == 1 { //V4=1
			fteid.IPv4Addr = getValue(data, 4)
		}
		if fteid.Flag&0b00000010 == 1 { //V6=1
			fteid.IPv6Addr = getValue(data, 16)
		}
	}
	return &fteid
}

func EncodeTSEID(fteid FTEID) []byte {
	return setValue(fteid.EType, fteid.ELength, fteid.Flag, fteid.TEID, fteid.IPv4Addr, fteid.IPv6Addr, fteid.ChooseID)
}

//判断是否含有FTEID
func HasTEID(fteid FTEID) bool {
	if fteid.EType == 0 {
		return false
	}
	return true
}
