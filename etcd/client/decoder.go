package client

import (
	"encoding/json"
	"fmt"
	"time"
)

type Decoder func(data []byte, config interface{}) error

// DefaultDecoder Default decoder
func DefaultDecoder(data []byte, config interface{}) error {
	return json.Unmarshal(data, config)
}

type EtcdConfig struct {
	ClientBasicInfo EndpointBasicInfo `yaml:"ClientBasicInfo"`
	HostPorts       int               `yaml:"hostports"`
	DestService     string            `yaml:"DestService"`
	Protocol        string            `yaml:"protocol"`
	Connection      Connection        `yaml:"Connection"`
}

func (c *EtcdConfig) ToString() string {
	return fmt.Sprintf("ClientBasicInfo: %v\n"+
		" HostPorts: %d\n"+
		" DestService: %s\n"+
		" Protocol: %s\n"+
		" Connection: %v\n",
		c.ClientBasicInfo, c.HostPorts, c.DestService, c.Protocol, c.Connection)
}

// EndpointBasicInfo should be immutable after created.
type EndpointBasicInfo struct {
	ServiceName string            `yaml:"ServiceName"`
	Method      string            `yaml:"Method"`
	Tags        map[string]string `yaml:"Tags"`
}

// IdleConfig contains idle configuration for long-connection pool.
type IdleConfig struct {
	MinIdlePerAddress int           `yaml:"MinIdlePerAddress"`
	MaxIdlePerAddress int           `yaml:"MaxIdlePerAddress"`
	MaxIdleGlobal     int           `yaml:"MaxIdleGlobal"`
	MaxIdleTimeout    time.Duration `yaml:"MaxIdleTimeout"`
}
type MuxConnection struct {
	ConnNum int `yaml:"connNum"`
}

type Connection struct {
	Method         string        `yaml:"method"`
	LongConnection IdleConfig    `yaml:"LongConnection"`
	MuxConnection  MuxConnection `yaml:"MuxConnection"`
}
