package config

import (
	"sync"
)

// ClientConfigDefn definition
type ClientConfigDefn struct {
	Addr string
	Port string
}

var cInstance *ClientConfigDefn
var cOnce sync.Once

// GetClientConfig get global client configuration
func GetClientConfig() *ClientConfigDefn {
	cOnce.Do(func() {
		cInstance = &ClientConfigDefn{
			Addr: "localhost",
			Port: DefaultPort,
		}
	})
	return cInstance
}
