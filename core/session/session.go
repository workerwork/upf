package session

import (
	"github.com/workerwork/upf/core/elem"
)

type Session struct {
	SessionDB //node节点数据库接口
	Elements
}

type Elements struct {
	LocalNodeID             elem.NodeID
	RemoteNodeID            elem.NodeID
	LocalRecoveryTimeStamp  elem.RecoveryTimeStamp
	RemoteRecoveryTimeStamp elem.RecoveryTimeStamp
}
