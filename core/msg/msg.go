package msg

import (
	"bytes"
	"encoding/binary"
	. "github.com/workerwork/upf/core/elem"
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
	Cause             //Cause
	NodeID            //NodeID
	RecoveryTimeStamp //RecoveryTimeStamp
	FSEID             //FSEID
	PDNType           //PDNType
	CreatePDR         //CreatePDR
	CreateFAR         //CreateFAR
	CreateURR         //CreateURR
	CreateQER         //CreateQER
}

//解析函数
func Parse(buf *bytes.Buffer) *Msg {
	var m Msg
	//解析消息头
	var b byte
	if err := binary.Read(buf, binary.BigEndian, &b); err != nil {
		log.Println(err) //TODO::
	}
	m.Version = b >> 5
	if b&0b00000010>>1 == 1 {
		m.MP = true
	}
	if b&0b00000001 == 1 {
		m.S = true
	}
	if err := binary.Read(buf, binary.BigEndian, &m.Type); err != nil {
		log.Println(err) //TODO::
	}
	if err := binary.Read(buf, binary.BigEndian, &m.Length); err != nil {
		log.Println(err) //TODO::
	}
	if m.S {
		if err := binary.Read(buf, binary.LittleEndian, &m.Head.SEID); err != nil {
			log.Println(err) //TODO::
		}
	}
	var b12 uint16
	if err := binary.Read(buf, binary.LittleEndian, &b12); err != nil {
		log.Println(err) //TODO::
	}
	var b3 uint8
	if err := binary.Read(buf, binary.LittleEndian, &b3); err != nil {
		log.Println(err) //TODO::
	}
	m.Sequence = uint32(b12<<8) | uint32(b3)
	var b4 byte
	if err := binary.Read(buf, binary.BigEndian, &b4); err != nil {
		log.Println(err) //TODO::
	}
	if m.MP {
		m.Priority = b4 >> 4
	}
	//解析消息体
	var cursor uint16
	dataLen := getDataLen(&m)
	for cursor < dataLen {
		var (
			eType IEType
			eLen  uint16
		)
		if err := binary.Read(buf, binary.BigEndian, &eType); err != nil {
			log.Println(err) //TODO::
		}
		if err := binary.Read(buf, binary.BigEndian, &eLen); err != nil {
			log.Println(err) //TODO::
		}
		e := make([]byte, eLen)
		//eValue := *bytes.NewBuffer(make([]byte, eLen))
		if err := binary.Read(buf, binary.BigEndian, e); err != nil {
			log.Println(err) //TODO::
		}
		eValue := bytes.NewBuffer(e)
		switch eType {
		case IETypeNodeID:
			m.NodeID = *DecodeNodeID(eValue, eLen)
		case IETypeRecoveryTimeStamp:
			m.RecoveryTimeStamp = *DecodeRecoveryTimeStamp(eValue, eLen)
		case IETypeUPFunctionFeatures:
			log.Println("IETypeUPFunctionFeatures")
			//TODO::
		case IETypeCPFunctionFeatures:
			log.Println("IETypeCPFunctionFeatures")
			//TODO::
		case IETypeUserPlaneIPResourceInformation:
			log.Println("IETypeUserPlaneIPResourceInformation")
			//TODO::
		case IETypeFSEID:
			m.FSEID = *DecodeFSEID(eValue, eLen)
		case IETypePDNType:
			m.PDNType = *DecodePDNType(eValue, eLen)
		case IETypeCreatePDR:
			m.CreatePDR = *DecodeCreatePDR(eValue, eLen)
		case IETypeCreateFAR:
			m.CreateFAR = *DecodeCreateFAR(eValue, eLen)
		case IETypeCreateURR:
			m.CreateURR = *DecodeCreateURR(eValue, eLen)
		case IETypeCreateQER:
			m.CreateQER = *DecodeCreateQER(eValue, eLen)
		default:
			log.Println("msg err: unknown tlv type", eType)
		}
		cursor = cursor + eLen + 4
	}
	return &m
}

//封装函数
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
		log.Println(err) //TODO::
	}
	if err := binary.Write(buf, binary.BigEndian, m.Type); err != nil {
		log.Println(err) //TODO::
	}
	if err := binary.Write(buf, binary.BigEndian, m.Length); err != nil {
		log.Println(err) //TODO::
	}
	if m.S {
		if err := binary.Write(buf, binary.BigEndian, m.Head.SEID); err != nil {
			log.Println(err) //TODO::
		}
	}
	b3 := make([]byte, 3)
	b3[0] = byte((m.Sequence >> 16) & 0xFF)
	b3[1] = byte((m.Sequence >> 8) & 0xFF)
	b3[2] = byte(m.Sequence & 0xFF)
	if err := binary.Write(buf, binary.BigEndian, b3); err != nil {
		log.Println(err) //TODO::
	}
	if m.MP {
		if err := binary.Write(buf, binary.BigEndian, m.Priority<<4); err != nil {
			log.Println(err) //TODO::
		}
	} else {
		if err := binary.Write(buf, binary.BigEndian, byte(0x00)); err != nil {
			log.Println(err) //TODO::
		}
	}
	//写入信元
	data := buf.Bytes()
	switch {
	case HasCause(m.Cause):
		SetValue(data, EncodeCause(m.Cause))
		fallthrough
	case HasNodeID(m.NodeID):
		SetValue(data, EncodeNodeID(m.NodeID))
		fallthrough
	case HasRecoveryTimeStamp(m.RecoveryTimeStamp):
		SetValue(data, EncodeRecoveryTimeStamp(m.RecoveryTimeStamp))
		fallthrough
	case HasFSEID(m.FSEID):
		SetValue(data, EncodeFSEID(m.FSEID))
		fallthrough
	case HasPDNType(m.PDNType):
		SetValue(data, EncodePDNType(m.PDNType))
		fallthrough
	case HasCreatePDR(m.CreatePDR):
		SetValue(data, EncodeCreatePDR(m.CreatePDR))
		fallthrough
	case HasCreateFAR(m.CreateFAR):
		SetValue(data, EncodeCreateFAR(m.CreateFAR))
		fallthrough
	case HasCreateURR(m.CreateURR):
		SetValue(data, EncodeCreateURR(m.CreateURR))
		fallthrough
	case HasCreateQER(m.CreateQER):
		SetValue(data, EncodeCreateQER(m.CreateQER))
	}
	log.Println("*********response: ", m)
	log.Printf("*********response: %0x\n", buf.Bytes())
	return bytes.NewBuffer(data)
}
