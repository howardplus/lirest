package source

import (
	log "github.com/Sirupsen/logrus"
	_ "github.com/howardplus/lirest/util"
	_ "os"
)

type SysctlExtractor struct {
	path string
}

func NewSysctlExtractor(path string) *SysctlExtractor {
	return &SysctlExtractor{path: path}
}

// implements the Extractor interface
func (e *SysctlExtractor) Extract(conv Converter) (interface{}, error) {
	log.Info("Run Sysctl extractor")
	return nil, nil
}
