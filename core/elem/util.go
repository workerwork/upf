package elem

import (
	"bytes"
	"encoding/binary"
	"log"
)

func getValue(data []byte, len uint16) []byte {
	buf := bytes.NewBuffer(data)
	value := make([]byte, len)
	if err := binary.Read(buf, binary.BigEndian, &value);err != nil {
		log.Println("err")	//TODO::
	}
	return value
}
