package node

import (
	"github.com/workerwork/upf/core/elem"
)

type CacheNode map[string]ElemInfo

func (cn CacheNode) Insert(elemInfo ElemInfo) error {
	cn[Bytes2Str(elemInfo.RemoteNodeID.NodeID)] = elemInfo
	return nil
}

func (cn CacheNode) Remove(nodeID elem.NodeID) error {
	delete(cn, Bytes2Str(nodeID.NodeID))
	return nil
}

func (cn CacheNode) IsExist(nodeID elem.NodeID) (bool, error) {
	if _, ok := cn[Bytes2Str(nodeID.NodeID)]; ok {
		return true, nil
	} else {
		return false, nil
	}
}

func (cn CacheNode) GetLocalNode(nodeID elem.NodeID) (elem.NodeID, error) {
	return cn[Bytes2Str(nodeID.NodeID)].LocalNodeID, nil
}

func (cn CacheNode) GetRemoteNode(nodeID elem.NodeID) (elem.NodeID, error) {
	return cn[Bytes2Str(nodeID.NodeID)].RemoteNodeID, nil
}

func (cn CacheNode) GetLocalRecoveryTimeStamp(nodeID elem.NodeID) (elem.RecoveryTimeStamp, error) {
	return cn[Bytes2Str(nodeID.NodeID)].LocalRecoveryTimeStamp, nil
}

func (cn CacheNode) GetRemoteRecoveryTimeStamp(nodeID elem.NodeID) (elem.RecoveryTimeStamp, error) {
	return cn[Bytes2Str(nodeID.NodeID)].RemoteRecoveryTimeStamp, nil
}
