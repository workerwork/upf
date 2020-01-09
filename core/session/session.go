package session

import (
	. "github.com/workerwork/upf/core/elem"
)

type Session struct {
	ID        string
	SessionDB //node节点数据库接口
	PDR       CreatePDR
	FAR       CreateFAR
	URR       CreateURR
	QER       CreateQER
	Precedence
}
