package core

import (
	"bytes"
	. "github.com/workerwork/upf/core/msg"
	. "github.com/workerwork/upf/core/node"
	"log"
)

func Run(buf *bytes.Buffer) {
	cn := NewCacheNode()
	m := Parse(buf)
	log.Println(m)
	switch m.Type {
	case PFCPMsgTypeAssociationSetupRequest:
		log.Println(cn.HandlePFCPMsgTypeAssociationSetupRequest(m).Pack().Bytes())
	case PFCPMsgTypeAssociationUpdateRequest:
		log.Println(cn.HandlePFCPMsgTypeAssociationUpdateRequest(m).Pack().Bytes())
	case PFCPMsgTypeAssociationReleaseRequest:
		log.Println(cn.HandlePFCPMsgTypeAssociationReleaseRequest(m).Pack().Bytes())
	case PFCPMsgTypeNodeReportResponse:
		//TODO::
	case PFCPMsgTypeSessionEstablishmentRequest:

	case PFCPMsgTypeSessionModificationRequest:

	case PFCPMsgTypeSessionDeletionRequest:

	default:
		//TODO::
	}

}
