package config

import (
	"sync"
)

const (
	DefaultPort = "8080" // the default port lirest listens on
)

// configuration definition
type ConfigDefn struct {
	Addr     string
	Port     string
	Security SecurityConfigDefn
	DescPath string
	Verbose  bool
	Quiet    bool
}

var instance *ConfigDefn
var once sync.Once

// The singleton config instance
func GetConfig() *ConfigDefn {
	once.Do(func() {
		instance = &ConfigDefn{
			Addr: "",
			Port: DefaultPort,
		}
	})
	return instance
}
