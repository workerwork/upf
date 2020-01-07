package node

import (
	. "github.com/workerwork/upf/core/elem"
	. "github.com/workerwork/upf/core/msg"
	"log"
)

type Node struct {
	NodeDB //node节点数据库接口
	Elements
}

type Elements struct {
	LocalNodeID             NodeID
	RemoteNodeID            NodeID
	LocalRecoveryTimeStamp  RecoveryTimeStamp
	RemoteRecoveryTimeStamp RecoveryTimeStamp
}

func NewCacheNode() *Node {
	return &Node{
		NodeDB: CacheNode{},
	}
}

func (node *Node) HandlePFCPMsgTypeAssociationSetupRequest(reqMsg *Msg) *Msg {
	respMsg := Msg{
		Head: Head{
			Version:  1,
			MP:       false,
			S:        false,
			Type:     PFCPMsgTypeAssociationSetupResponse,
			SEID:     0,
			Sequence: reqMsg.Sequence,
			Priority: reqMsg.Priority,
		},
		NodeID:            *NewIPv4NodeID([]byte{1, 1, 1, 1}),        //从配置读取	//TODO::
		RecoveryTimeStamp: *NewRecoveryTimeStamp([]byte{2, 2, 2, 2}), //从配置读取	//TODO::
		Cause:             *NewCause(CauseSuccess),
	}
	//信元类型２字节＋信元长度表示２字节＋消息头sequence3字节+space1字节
	respMsg.Length = respMsg.NodeID.ELength + 4 + respMsg.RecoveryTimeStamp.ELength + 4 + respMsg.Cause.ELength + 4 + 4
	if !HasNodeID(reqMsg.NodeID) || !HasRecoveryTimeStamp(reqMsg.RecoveryTimeStamp) {
		log.Println("false") //TODO::
		respMsg.Cause = Cause{Cause: CauseMandatoryIEMissing}
	}
	if isExist, _ := node.NodeDB.IsExist(reqMsg.NodeID); isExist {
		log.Println("Node is exist!") //TODO::
		respMsg.Cause = Cause{Cause: CauseUnspecifiedReason}
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

func (node *Node) HandlePFCPMsgTypeAssociationUpdateRequest(reqMsg *Msg) *Msg {
	respMsg := Msg{
		Head: Head{
			Version:  1,
			MP:       false,
			S:        false,
			Type:     PFCPMsgTypeAssociationUpdateResponse,
			SEID:     0,
			Sequence: reqMsg.Sequence,
			Priority: reqMsg.Priority,
		},
		NodeID:            *NewIPv4NodeID([]byte{1, 1, 1, 1}),        //从配置读取	//TODO::
		Cause:             *NewCause(CauseSuccess),
	}
	respMsg.Length = respMsg.NodeID.ELength + 4 + respMsg.Cause.ELength + 4 + 4
	if !HasNodeID(reqMsg.NodeID) {
		log.Println("false") //TODO::
		respMsg.Cause = Cause{Cause: CauseMandatoryIEMissing}
	}
	if isExist, _ := node.NodeDB.IsExist(reqMsg.NodeID); !isExist {
		log.Println("Node is not exist!") //TODO::
		respMsg.Cause = Cause{Cause: CauseUnspecifiedReason}
	} else {
		//更新node节点
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

func (node *Node) HandlePFCPMsgTypeAssociationReleaseRequest(reqMsg *Msg) *Msg {
	respMsg := Msg{
		Head: Head{
			Version:  1,
			MP:       false,
			S:        false,
			Type:     PFCPMsgTypeAssociationReleaseResponse,
			SEID:     0,
			Sequence: reqMsg.Sequence,
			Priority: reqMsg.Priority,
		},
		NodeID:            *NewIPv4NodeID([]byte{1, 1, 1, 1}),        //从配置读取	//TODO::
		Cause:             *NewCause(CauseSuccess),
	}
	respMsg.Length = respMsg.NodeID.ELength + 4 + respMsg.Cause.ELength + 4 + 4
	if !HasNodeID(reqMsg.NodeID) {
		log.Println("false") //TODO::
		respMsg.Cause = Cause{Cause: CauseMandatoryIEMissing}
	}
	if isExist, _ := node.NodeDB.IsExist(reqMsg.NodeID); !isExist {
		log.Println("Node is not exist!") //TODO::
		respMsg.Cause = Cause{Cause: CauseUnspecifiedReason}
	} else {
		//删除node节点
		if err := node.NodeDB.Remove(reqMsg.NodeID); err != nil {
			//TODO::
		}
	}
	return &respMsg
}

func (node *Node) HandlePFCPMsgTypeNodeReportResponse(reqMsg *Msg) *Msg {
	return &Msg{}
}
