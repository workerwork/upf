package elem

import (
	"bytes"
	"encoding/binary"
	"log"
)

func getValue(buf *bytes.Buffer, len uint16) []byte {
	value := make([]byte, len)
	if err := binary.Read(buf, binary.BigEndian, &value); err != nil {
		log.Println(buf, value, len, err) //TODO::
	}
	return value
}
func SetValue(data ...interface{}) *bytes.Buffer {
	buf := bytes.NewBuffer([]byte{})
	for _, d := range data {
		if err := binary.Write(buf, binary.BigEndian, d); err != nil {
			log.Println(err) //TODO::
		}
	}
	return buf
}
