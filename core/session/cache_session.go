package session

type SessionInfo struct {
	NodeInfo   node.NodeInfo
	LocalSEID  []byte
	RemoteSEID []byte
}

type CacheSession map[string]SessionIfo

func (cs Session) Insert(s Session) {

}
