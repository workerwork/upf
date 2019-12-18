package elem

import (
	"bytes"
	"encoding/binary"
)

func getValue(data []byte, len uint16) []byte {
	buf := bytes.NewBuffer(data)
	value := make([]byte, len)
	binary.Read(buf, binary.BigEndian, &value)
	return value
}

func setBuffer(t IEType, l uint16, v []byte) []byte {
	buf := bytes.NewBuffer([]byte{})
	binary.Write(buf, binary.BigEndian, t)
	binary.Write(buf, binary.BigEndian, l)
	binary.Write(buf, binary.BigEndian, v)
	return buf.Bytes()
}

func setValue(t byte, v []byte) []byte {
	buf := bytes.NewBuffer([]byte{})
	binary.Write(buf, binary.BigEndian, t)
	binary.Write(buf, binary.BigEndian, v)
	return buf.Bytes()
}
