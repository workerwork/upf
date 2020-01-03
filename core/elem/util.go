package elem

import (
	"bytes"
	"encoding/binary"
	"log"
)

/*
func getValue(data []byte, len uint16) []byte {
	buf := bytes.NewBuffer(data)
	value := make([]byte, len)
	if err := binary.Read(buf, binary.BigEndian, &value); err != nil {
		log.Println("err") //TODO::
	}
	return value
}*/

func getValue(buf *bytes.Buffer, len uint16) []byte {
	value := make([]byte, len)
	if err := binary.Read(buf, binary.BigEndian, &value); err != nil {
		log.Println("err") //TODO::
	}
	return value
}

func setValue(data ...interface{}) []byte {
	buf := bytes.NewBuffer([]byte{})
	for _, d := range data {
		if err := binary.Write(buf, binary.BigEndian, d); err != nil {
			log.Println(err) //TODO::
		}
	}
	return buf.Bytes()
}
