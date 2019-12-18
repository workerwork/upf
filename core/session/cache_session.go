package session

import (
	"N4test/core/elem"
	"bytes"
	"encoding/binary"
	"log"
)

type SessionInfo struct {
	NodeInfo   node.NodeInfo
	LocalSEID  []byte
	RemoteSEID []byte
}

type CacheSession map[string]SessionIfo

func (cs Session) Insert(s Session) {

}
