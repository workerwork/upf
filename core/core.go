package core

import (
	"bytes"
	"github.com/workerwork/upf/core/msg"
	"github.com/workerwork/upf/core/node"
	"log"
)

func Run(buf *bytes.Buffer){
	m := msg.Parse(buf)
	log.Println(m)
	switch m.Type {
	case msg.PFCPMsgTypeAssociationSetupRequest:
		cn := node.NewCacheNode()
		log.Println(cn.HandlePFCPMsgTypeAssociationSetupRequest(m).Pack().Bytes())
	case msg.PFCPMsgTypeAssociationUpdateRequest:
	case msg.PFCPMsgTypeAssociationReleaseRequest:
	case msg.PFCPMsgTypeNodeReportResponse:
	default:
		//TODO::
	}

}