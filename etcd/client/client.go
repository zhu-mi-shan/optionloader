package client

type Client interface {
	NewReader(etcdPath string) *EtcdReader
}

type EtcdClient struct {
}
