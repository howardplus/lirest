package inject

import (
	log "github.com/Sirupsen/logrus"
	"github.com/howardplus/lirest/describe"
	"github.com/howardplus/lirest/util"
	"io/ioutil"
)

// PathInjector defines a path injector based on write format
type PathInjector struct {
	format describe.DescriptionWriteFormat
	vars   map[string]string
}

// NewPathInjector creates a new path injector
func NewPathInjector(format describe.DescriptionWriteFormat, vars map[string]string) *PathInjector {
	return &PathInjector{
		format: format,
		vars:   vars,
	}
}

// Inject injects data to the defined path
func (inj *PathInjector) Inject(data string) error {

	log.WithFields(log.Fields{
		"data": data,
		"path": inj.format.Path,
	}).Info("Inject data")

	// validate data first
	if err := NewValidator(inj.format).Validate(data); err != nil {
		return err
	}

	// TODO: read old value
	old := ""

	// send data to path
	buf := []byte(data)
	if err := ioutil.WriteFile(inj.format.Path, buf, 0400); err != nil {
		log.Error(err.Error())
		return util.NewError(err.Error())
	}

	// send successful, record this job
	RecordJob(inj, old, data)
	return nil
}

// Name returns the name of the injector
// for PathInjector, the path is the name
func (inj *PathInjector) Name() string {
	return "path:" + inj.format.Path
}
