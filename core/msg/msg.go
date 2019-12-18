package msg

import (
	"github.com/workerwork/upf/core/elem"
	"bytes"
	"encoding/binary"
	"log"
)

type Head struct {
	Version  byte //100
	MP       bool
	S        bool
	Type     PFCPMsgType
	Length   uint16 //indicate the length of the message in octets excluding the mantory 4 octets
	SEID     uint64
	Sequence uint32 //3byte
	Priority byte
}

type Msg struct {
	Head
	Cause             elem.Cause
	NodeID            elem.NodeID
	RecoveryTimeStamp elem.RecoveryTimeStamp
}

func Parse(buf *bytes.Buffer) *Msg {
	var m Msg
	var b byte
	if err := binary.Read(buf, binary.BigEndian, &b); err != nil {
		log.Println("err") //TODO::
	}
	m.Version = b >> 5
	if b&0b00000010 == 1 {
		m.MP = true
	}
	if b&0b00000001 == 1 {
		m.S = true
	}
	if err := binary.Read(buf, binary.BigEndian, &m.Type); err != nil {
		log.Println("err") //TODO::
	}
	if err := binary.Read(buf, binary.BigEndian, &m.Length); err != nil {
		log.Println("err") //TODO::
	}
	if m.S {
		if err := binary.Read(buf, binary.LittleEndian, &m.SEID); err != nil {
			log.Println("err") //TODO::
		}
	}
	var b12 uint16
	if err := binary.Read(buf, binary.LittleEndian, &b12); err != nil {
		log.Println("err") //TODO::
	}
	var b3 uint8
	if err := binary.Read(buf, binary.LittleEndian, &b3); err != nil {
		log.Println("err") //TODO::
	}
	m.Sequence = uint32(b12<<8) | uint32(b3)
	var b4 byte
	if err := binary.Read(buf, binary.BigEndian, &b4); err != nil {
		log.Println("err") //TODO::
	}
	if m.MP {
		m.Priority = b4 >> 4
	}
	//解析消息体
	var cursor uint16
	dataLen := getDataLen(&m)
	for cursor < dataLen {
		var (
			eType elem.IEType
			eLen  uint16
		)
		if err := binary.Read(buf, binary.BigEndian, &eType); err != nil {
			log.Println("err") //TODO::
		}
		if err := binary.Read(buf, binary.BigEndian, &eLen); err != nil {
			log.Println("err") //TODO::
		}
		eValue := make([]byte, eLen)
		if err := binary.Read(buf, binary.BigEndian, &eValue); err != nil {
			log.Println("err") //TODO::
		}
		switch eType {
		case elem.IETypeNodeID:
			m.NodeID = *elem.DecodeNodeID(eValue, eLen)
		case elem.IETypeRecoveryTimeStamp:
			m.RecoveryTimeStamp = *elem.DecodeRecoveryTimeStamp(eValue, eLen)

		case elem.IETypeUPFunctionFeatures:
			//TODO::

		case elem.IETypeCPFunctionFeatures:
			//TODO::
		case elem.IETypeUserPlaneIPResourceInformation:
			//TODO::

		default:
			log.Println("err: unknown tlv type", eType)
		}
		cursor = cursor + eLen + 4
	}
	return &m
}

func (m *Msg) Pack() *bytes.Buffer {
	buf := bytes.NewBuffer([]byte{})
	b := m.Version << 5
	if m.MP {
		b = b | 0b00000010
	}
	if m.S {
		b = b | 0b00000001
	}
	if err := binary.Write(buf, binary.BigEndian, b); err != nil {
		log.Println("err") //TODO::
	}
	if err := binary.Write(buf, binary.BigEndian, m.Type); err != nil {
		log.Println("err") //TODO::
	}
	if err := binary.Write(buf, binary.BigEndian, m.Length); err != nil {
		log.Println("err") //TODO::
	}
	if m.S {
		if err := binary.Write(buf, binary.BigEndian, m.SEID); err != nil {
			log.Println("err") //TODO::
		}
	}
	b3 := make([]byte, 3)
	b3[0] = byte((m.Sequence >> 16) & 0xFF)
	b3[1] = byte((m.Sequence >> 8) & 0xFF)
	b3[2] = byte(m.Sequence & 0xFF)
	if err := binary.Write(buf, binary.BigEndian, b3); err != nil {
		log.Println("err") //TODO::
	}
	if m.MP {
		if err := binary.Write(buf, binary.BigEndian, m.Priority<<4); err != nil {
			log.Println("err") //TODO::
		}
	} else {
		if err := binary.Write(buf, binary.BigEndian, byte(0x00)); err != nil {
			log.Println("err") //TODO::
		}
	}
	//TODO::判断具体的信元是否存在，如果存在则写入
	if m.NodeID.EType != 0 {
		if err := binary.Write(buf, binary.BigEndian, elem.EncodeNodeID(m.NodeID)); err != nil {
			log.Println("err") //TODO::
		}
	}
	if m.RecoveryTimeStamp.EType != 0 {
		if err := binary.Write(buf, binary.BigEndian, elem.EncodeRecoveryTimeStamp(m.RecoveryTimeStamp)); err != nil {
			log.Println("err") //TODO::
		}
	}
	log.Println("*********response: ", m)
	log.Printf("*********response: %0x\n", buf.Bytes())
	return buf
}
