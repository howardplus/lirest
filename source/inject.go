package source

import (
	log "github.com/Sirupsen/logrus"
	"github.com/howardplus/lirest/describe"
	"github.com/howardplus/lirest/inject"
	"github.com/howardplus/lirest/util"
	"io/ioutil"
)

// NewInjector creates an injector based on the source type
func NewInjector(s describe.DescriptionSource, f describe.DescriptionFormat) (inject.Injector, error) {
	var injector inject.Injector

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

// GenericInjector
type GenericInjector struct {
	path   string
	format describe.DescriptionFormat
}

// NewGenericInjector
func NewGenericInjector(path string, format describe.DescriptionFormat) *GenericInjector {
	return &GenericInjector{path: path, format: format}
}

func (inj *GenericInjector) Inject(data string) error {

	log.WithFields(log.Fields{
		"data": data,
		"path": inj.path,
	}).Info("Inject data")

	// validate data first
	if err := validate(data, inj.format); err != nil {
		return err
	}

	// read old value
	old := ""

	// send data to file
	buf := []byte(data)
	if err := ioutil.WriteFile(inj.path, buf, 0400); err != nil {
		log.Error(err.Error())
		return util.NewError(err.Error())
	}

	// send successful, record this job
	inject.RecordJob(inj, old, data)
	return nil
}

func (inj *GenericInjector) Name() string {
	return inj.path
}

func validate(data string, format describe.DescriptionFormat) error {
	// TODO: need to validate
	return nil
}
