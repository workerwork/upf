package elem

import (
	"bytes"
)

type UEIPAddressFlag byte

type UEIPAddress struct {
	EType      IEType
	ELength    uint16
	Flag       UEIPAddressFlag
	IPv4Addr   []byte //4bytes
	IPv6Addr   []byte //16bytes
	IPv6Prefix byte
}

func DecodeUEIPAddress(buf *bytes.Buffer, len uint16) *UEIPAddress {
	u := UEIPAddress{
		EType:   IETypeUEIPAddress,
		ELength: len,
		Flag:    UEIPAddressFlag(GetValue(buf, 1)[0]),
	}
	if u.Flag&0b00010000 != 0 { //CH=1
		//TODO::
	} else {
		switch {
		case u.Flag&0b00000010 != 0: //v4=1
			u.IPv4Addr = GetValue(buf, 4)
			fallthrough
		case u.Flag&0b00000001 != 0: //v6=1
			u.IPv6Addr = GetValue(buf, 16)
			fallthrough
		case u.Flag&0b00000100 != 0: //S/D=1
			fallthrough
		case u.Flag&0b00001000 != 0: //ipv6D=1
			u.IPv6Prefix = GetValue(buf, 1)[0]
		}
	}
	return &u
}

func EncodeUEIPAddress(u UEIPAddress) []byte {
	ret := SetValue(u.EType, u.ELength, u.Flag, u.IPv4Addr, u.IPv6Addr)
	if u.Flag&0b00010000>>4 == 0 && u.Flag&0b00001000>>3 == 1 {
		SetValue(ret, u.IPv6Prefix)
	}
	return ret
}

//判断是否含有UEIPAddress
func HasUEIPAddress(u UEIPAddress) bool {
	if u.EType == 0 {
		return false
	}
	return true
}
