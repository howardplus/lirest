package inject

import (
	log "github.com/Sirupsen/logrus"
	"github.com/howardplus/lirest/describe"
	"github.com/howardplus/lirest/util"
	"io/ioutil"
)

// Injector implements the Inject method
// it validate and writes the data into the "source"
type Injector interface {
	Inject(data string) error
	Name() string
}

// NewInjector creates an injector based on the source type
func NewInjector(s describe.DescriptionSource, f describe.DescriptionWriteFormat) (Injector, error) {
	var injector Injector

	switch s.Type {
	case "procfs", "sysfs", "sysctl":
		injector = NewGenericInjector(s.Path, f)
	}

	// found an injector, use it
	if injector != nil {
		return injector, nil
	}

	return nil, util.NewError("Internal error: unknown input type")
}

// GenericInjector defines a generic injector based on write format
type GenericInjector struct {
	path   string
	format describe.DescriptionWriteFormat
}

// NewGenericInjector creates a new generic injector
func NewGenericInjector(path string, format describe.DescriptionWriteFormat) *GenericInjector {
	return &GenericInjector{
		path:   path,
		format: format,
	}
}

// Inject injects data
func (inj *GenericInjector) Inject(data string) error {

	log.WithFields(log.Fields{
		"data": data,
		"path": inj.path,
	}).Info("Inject data")

	// validate data first
	if err := NewValidator(inj.format).Validate(data); err != nil {
		return err
	}

	// TODO: read old value
	old := ""

	// send data to file
	buf := []byte(data)
	if err := ioutil.WriteFile(inj.path, buf, 0400); err != nil {
		log.Error(err.Error())
		return util.NewError(err.Error())
	}

	// send successful, record this job
	RecordJob(inj, old, data)
	return nil
}

// Name returns the name of the injector
func (inj *GenericInjector) Name() string {
	return inj.path
}
