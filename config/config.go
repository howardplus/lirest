package config

import (
	"sync"
)

const (
	// the project name
	ProjectName = "LiREST"
	// the default port lirest listens on
	DefaultPort = "8080"
)

// ConfigDefn definition
type ConfigDefn struct {
	Addr     string
	Port     string
	DescPath string
	DescUrl  string
	NoSysctl bool
	Verbose  bool
	Quiet    bool
	Watch    bool // watch for changes in description files
	Pretty   bool
}

var instance *ConfigDefn
var once sync.Once

// GetConfig returns the singleton config instance
func GetConfig() *ConfigDefn {
	once.Do(func() {
		instance = &ConfigDefn{
			Addr: "",
			Port: DefaultPort,
		}
	})
	return instance
}
