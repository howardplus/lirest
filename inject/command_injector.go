package inject

import (
	log "github.com/Sirupsen/logrus"
	"github.com/howardplus/lirest/describe"
	"github.com/howardplus/lirest/util"
	"os/exec"
)

// CommandInjector defines a command injector based on write format
type CommandInjector struct {
	format describe.DescriptionWriteFormat
}

// NewCommandInjector creates a new path injector
func NewCommandInjector(format describe.DescriptionWriteFormat) *CommandInjector {
	return &CommandInjector{
		format: format,
	}
}

// Inject injects data using the defined command
func (inj *CommandInjector) Inject(data string) error {

	log.WithFields(log.Fields{
		"data":    data,
		"command": inj.format.Command,
	}).Info("Inject data")

	// validate data first
	if err := NewValidator(inj.format).Validate(data); err != nil {
		return err
	}

	// send data to command
	vars := make(map[string]string, 1)
	if data != "" {
		vars["data"] = data
	}

	if cmdStr, err := util.FillVars(inj.format.Command, vars); err != nil {
		return err
	} else {
		// issue command
		log.Info(cmdStr)
		cmd := exec.Command("sh", "-c", cmdStr)
		if err := cmd.Run(); err != nil {
			return err
		}
	}

	// send successful, record this job
	RecordJob(inj, "", data)
	return nil
}

// Name returns the name of the injector
func (inj *CommandInjector) Name() string {
	return "command"
}
