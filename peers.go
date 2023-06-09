package mccache

// PeerPicker是必须实现的接口，用于选择给定键的Peer
type PeerPicker interface {
	PickPeer(key string) (peer PeerGetter, ok bool)
}

// PeerGetter是Peer必须实现的接口，用于从对应的group中查找缓存值
type PeerGetter interface {
	Get(group string, key string) ([]byte, error)
}
