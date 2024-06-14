package client

type Reader interface {
	SetDecoder(decoder Decoder) error
	ReadToConfig(data []byte) error
	GetConfig(key string) (map[string]interface{}, error)
}

// EtcdReader yml/client/reader.go
type EtcdReader struct {
	config  EtcdConfig //配置文件读出结果
	decoder Decoder    //配置文件解码器
}
