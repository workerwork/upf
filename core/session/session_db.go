package session

import (
	"github.com/workerwork/upf/core/elem"
)

type SessionDB interface {
	Insert(ElemInfo) error
	Remove(elem.NodeID) error
	IsExist(elem.NodeID) (bool, error)
	GetLocalNode(elem.NodeID) (elem.NodeID, error)
	GetRemoteNode(elem.NodeID) (elem.NodeID, error)
	GetLocalRecoveryTimeStamp(elem.NodeID) (elem.RecoveryTimeStamp, error)
	GetRemoteRecoveryTimeStamp(elem.NodeID) (elem.RecoveryTimeStamp, error)
}
