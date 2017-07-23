package inject

import (
	"github.com/howardplus/lirest/describe"
	"github.com/howardplus/lirest/util"
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
		injector = NewPathInjector(f)
	case "command":
		injector = NewCommandInjector(f)
	}

	// found an injector, use it
	if injector != nil {
		return injector, nil
	}

	return nil, util.NewError("Internal error: unknown input type")
}
