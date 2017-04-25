package source

import (
	log "github.com/Sirupsen/logrus"
	"github.com/howardplus/lirest/describe"
	"github.com/howardplus/lirest/sandbox"
	"github.com/howardplus/lirest/util"
	"io/ioutil"
	"time"
)

// Injector implements the Inject method
// it validate and writes the data into the "source"
type Injector interface {
	Inject(data string) (*sandbox.Job, error)
}

// NewInjector creates an injector based on the source type
func NewInjector(s describe.DescriptionSource, f describe.DescriptionFormat) (Injector, error) {
	var injector Injector

	switch s.Type {
	case "procfs":
	case "sysfs":
	case "sysctl":
		injector = NewGenericInjector(s.Path, f)
	}

	// found an injector, use it
	if injector != nil {
		return injector, nil
	}

	return nil, util.NewError("Internal error: unknown input type")
}

// GenericInjector
type GenericInjector struct {
	path   string
	format describe.DescriptionFormat
}

// NewGenericInjector
func NewGenericInjector(path string, format describe.DescriptionFormat) *GenericInjector {
	return &GenericInjector{path: path, format: format}
}

func (inj *GenericInjector) Inject(data string) (*sandbox.Job, error) {

	log.WithFields(log.Fields{
		"data": data,
		"path": inj.path,
	}).Info("Inject data")

	// validate data first
	if err := validate(data, inj.format); err != nil {
		return nil, err
	}

	// send data to file
	buf := []byte(data)
	if err := ioutil.WriteFile(inj.path, buf, 0400); err != nil {
		log.Error(err.Error())
		return nil, util.NewError(err.Error())
	}

	return &sandbox.Job{Time: time.Now(), Id: 0}, nil
}

func validate(data string, format describe.DescriptionFormat) error {
	// TODO: need to validate
	return nil
}
