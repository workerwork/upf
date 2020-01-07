package node

import (
	. "github.com/workerwork/upf/core/elem"
)

type CacheNode map[string]Elements

func (cn CacheNode) Insert(elemInfo Elements) error {
	cn[Bytes2Str(elemInfo.RemoteNodeID.NodeID)] = elemInfo
	return nil
}

func (cn CacheNode) Remove(nodeID NodeID) error {
	delete(cn, Bytes2Str(nodeID.NodeID))
	return nil
}

func (cn CacheNode) IsExist(nodeID NodeID) (bool, error) {
	if _, ok := cn[Bytes2Str(nodeID.NodeID)]; ok {
		return true, nil
	} else {
		return false, nil
	}
}

func (cn CacheNode) GetLocalNode(nodeID NodeID) (NodeID, error) {
	return cn[Bytes2Str(nodeID.NodeID)].LocalNodeID, nil
}

func (cn CacheNode) GetRemoteNode(nodeID NodeID) (NodeID, error) {
	return cn[Bytes2Str(nodeID.NodeID)].RemoteNodeID, nil
}

func (cn CacheNode) GetLocalRecoveryTimeStamp(nodeID NodeID) (RecoveryTimeStamp, error) {
	return cn[Bytes2Str(nodeID.NodeID)].LocalRecoveryTimeStamp, nil
}

func (cn CacheNode) GetRemoteRecoveryTimeStamp(nodeID NodeID) (RecoveryTimeStamp, error) {
	return cn[Bytes2Str(nodeID.NodeID)].RemoteRecoveryTimeStamp, nil
}
