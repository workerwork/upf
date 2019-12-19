package node

import (
	"github.com/workerwork/upf/core/elem"
	"github.com/workerwork/upf/core/msg"
	"log"
)

type Node struct {
	NodeDB //node节点数据库接口
	Elements
}

type Elements struct {
	LocalNodeID             elem.NodeID
	RemoteNodeID            elem.NodeID
	LocalRecoveryTimeStamp  elem.RecoveryTimeStamp
	RemoteRecoveryTimeStamp elem.RecoveryTimeStamp
}

func NewCacheNode() *Node {
	return &Node{
		NodeDB: CacheNode{},
	}
}

/**
func (node *Node) Run(reqMsg *msg.Msg) *msg.Msg {
	switch reqMsg.Type {
	case msg.PFCPMsgTypeAssociationSetupRequest:
		respMsg := msg.Msg{
			Head: msg.Head{
				Version:  1,
				MP:       false,
				S:        false,
				Type:     msg.PFCPMsgTypeAssociationSetupResponse,
				SEID:     0,
				Sequence: reqMsg.Sequence,
				Priority: reqMsg.Priority,
			},
			NodeID:            *elem.NewIPv4NodeID([]byte{1, 1, 1, 1}),        //从配置读取	//TODO::
			RecoveryTimeStamp: *elem.NewRecoveryTimeStamp([]byte{2, 2, 2, 2}), //从配置读取	//TODO::
			Cause:             *elem.NewCause(elem.CauseSuccess),
		}
		//信元类型２字节＋信元长度表示２字节＋消息头sequence3字节+space1字节
		respMsg.Length = respMsg.NodeID.ELength + 4 + respMsg.RecoveryTimeStamp.ELength + 4 + respMsg.Cause.ELength + 4 + 4
		if !elem.HasNodeID(reqMsg.NodeID) || !elem.HasRecoveryTimeStamp(reqMsg.RecoveryTimeStamp) {
			log.Println("false") //TODO::
			respMsg.Cause = elem.Cause{Cause: elem.CauseMandatoryIEMissing}
		}
		if isExist, _ := node.NodeDB.IsExist(reqMsg.NodeID); isExist {
			log.Println("Node is exist!") //TODO::
			respMsg.Cause = elem.Cause{Cause: elem.CauseUnspecifiedReason}
		} else {
			//添加node节点
			node.Elements = Elements{
				LocalNodeID:             respMsg.NodeID,
				RemoteNodeID:            reqMsg.NodeID,
				LocalRecoveryTimeStamp:  respMsg.RecoveryTimeStamp,
				RemoteRecoveryTimeStamp: reqMsg.RecoveryTimeStamp,
			}
			if err := node.NodeDB.Insert(node.Elements); err != nil {
				//TODO::
			}
		}
		return &respMsg
	case msg.PFCPMsgTypeAssociationUpdateRequest:
	case msg.PFCPMsgTypeAssociationReleaseRequest:
	case msg.PFCPMsgTypeHeartbeatRequest:
	case msg.PFCPMsgTypeNodeReportResponse:
	default:
		log.Println("unknown message!") //TODO::
	}
	return &msg.Msg{}
}*/

func (node *Node) HandlePFCPMsgTypeAssociationSetupRequest(reqMsg *msg.Msg) *msg.Msg {
	respMsg := msg.Msg{
		Head: msg.Head{
			Version:  1,
			MP:       false,
			S:        false,
			Type:     msg.PFCPMsgTypeAssociationSetupResponse,
			SEID:     0,
			Sequence: reqMsg.Sequence,
			Priority: reqMsg.Priority,
		},
		NodeID:            *elem.NewIPv4NodeID([]byte{1, 1, 1, 1}),        //从配置读取	//TODO::
		RecoveryTimeStamp: *elem.NewRecoveryTimeStamp([]byte{2, 2, 2, 2}), //从配置读取	//TODO::
		Cause:             *elem.NewCause(elem.CauseSuccess),
	}
	//信元类型２字节＋信元长度表示２字节＋消息头sequence3字节+space1字节
	respMsg.Length = respMsg.NodeID.ELength + 4 + respMsg.RecoveryTimeStamp.ELength + 4 + respMsg.Cause.ELength + 4 + 4
	if !elem.HasNodeID(reqMsg.NodeID) || !elem.HasRecoveryTimeStamp(reqMsg.RecoveryTimeStamp) {
		log.Println("false") //TODO::
		respMsg.Cause = elem.Cause{Cause: elem.CauseMandatoryIEMissing}
	}
	if isExist, _ := node.NodeDB.IsExist(reqMsg.NodeID); isExist {
		log.Println("Node is exist!") //TODO::
		respMsg.Cause = elem.Cause{Cause: elem.CauseUnspecifiedReason}
	} else {
		//添加node节点
		node.Elements = Elements{
			LocalNodeID:             respMsg.NodeID,
			RemoteNodeID:            reqMsg.NodeID,
			LocalRecoveryTimeStamp:  respMsg.RecoveryTimeStamp,
			RemoteRecoveryTimeStamp: reqMsg.RecoveryTimeStamp,
		}
		if err := node.NodeDB.Insert(node.Elements); err != nil {
			//TODO::
		}
	}
	return &respMsg
}

func (node *Node) HandlePFCPMsgTypeAssociationUpdateRequest(reqMsg *msg.Msg) *msg.Msg {}

func (node *Node) HandlePFCPMsgTypeAssociationReleaseRequest(reqMsg *msg.Msg) *msg.Msg {}

func (node *Node) HandlePFCPMsgTypeNodeReportResponse(reqMsg *msg.Msg) *msg.Msg {}
