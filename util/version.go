package util

import (
	"sync"
)

// VersionTag contains major and minor versions
type VersionTag struct {
	Major int
	Minor int
}

var instance *VersionTag
var once sync.Once

// The singleton version instance
func GetVersion() *VersionTag {
	once.Do(func() {
		instance = &VersionTag{
			Major: 0,
			Minor: 1,
		}
	})
	return instance
}
