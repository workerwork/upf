package node

import (
	. "github.com/workerwork/upf/core/elem"
)

type NodeDB interface {
	Insert(Elements) error
	Remove(NodeID) error
	IsExist(NodeID) (bool, error)
	GetLocalNode(NodeID) (NodeID, error)
	GetRemoteNode(NodeID) (NodeID, error)
	GetLocalRecoveryTimeStamp(NodeID) (RecoveryTimeStamp, error)
	GetRemoteRecoveryTimeStamp(NodeID) (RecoveryTimeStamp, error)
}

type SessionDB interface {
	Insert(Elements) error
	Remove(NodeID) error
	IsExist(NodeID) (bool, error)
	GetLocalNode(NodeID) (NodeID, error)
	GetRemoteNode(NodeID) (NodeID, error)
	GetLocalRecoveryTimeStamp(NodeID) (RecoveryTimeStamp, error)
	GetRemoteRecoveryTimeStamp(NodeID) (RecoveryTimeStamp, error)
}
