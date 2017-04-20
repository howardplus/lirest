package source

import (
	log "github.com/Sirupsen/logrus"
)

// SysctlExtractor
type SysctlExtractor struct {
	path string
}

// NewSysctlExtractor
func NewSysctlExtractor(path string) *SysctlExtractor {
	return &SysctlExtractor{path: path}
}

// Extract
// implements the Extractor interface
func (e *SysctlExtractor) Extract(conv Converter) (interface{}, error) {
	log.Info("Run Sysctl extractor")
	return nil, nil
}
