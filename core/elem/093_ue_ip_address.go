package elem

import "bytes"

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
		Flag:    UEIPAddressFlag(getValue(buf, 1)[0]),
	}
	if u.Flag&0b00010000>>4 == 1 { //CH=1
		//TODO::
	} else {
		if u.Flag&0b00000001 == 1 { //v6=1
			u.IPv6Addr = getValue(buf, 16)
		}
		if u.Flag&0b00000010>>1 == 1 { //v4=1
			u.IPv4Addr = getValue(buf, 4)
		}
		if u.Flag&0b00000100>>2 == 1 { //v4=1
			//TODO::
		}
		if u.Flag&0b00001000>>3 == 1 { //v4=1
			u.IPv6Prefix = getValue(buf, 1)[0]
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
