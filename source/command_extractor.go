package source

import (
	"bytes"
	log "github.com/Sirupsen/logrus"
	"github.com/howardplus/lirest/util"
	"os/exec"
	"strings"
	"time"
)

// CommandExtractor extracts data by running a command
type CommandExtractor struct {
	cmd  string
	conv Converter
	vars map[string]string
}

// NewCommandExtractor creates a command extractor
// that extract data from running a system command
func NewCommandExtractor(cmd string, conv Converter, vars map[string]string) *CommandExtractor {
	return &CommandExtractor{cmd: cmd, conv: conv, vars: vars}
}

// Extract extracts data by running a command
func (e *CommandExtractor) Extract() (*ExtractOutput, error) {
	log.WithFields(log.Fields{
		"cmd":  e.cmd,
		"vars": e.vars,
	}).Debug("Extract from command")

	// create command from variables
	cmd, err := util.FillVars(e.cmd, e.vars)
	if err != nil {
		return nil, util.NewError("Failed to generate command")
	}

	// run command
	cmdTokens := strings.Split(cmd, " ")
	out, _ := exec.Command(cmdTokens[0], cmdTokens[1:]...).Output()
	buf := bytes.NewBuffer(out)

	// give it to the converter
	data, err := e.conv.ConvertStream(buf)
	if err != nil {
		return nil, err
	}

	log.WithFields(log.Fields{
		"cmd": cmd,
	}).Debug("Convert successful")

	return &ExtractOutput{
		Name: e.conv.Name(),
		Time: time.Now(),
		Data: data,
	}, nil
}
