package inject

import (
	"github.com/howardplus/lirest/job"
)

// Injector implements the Inject method
// it validate and writes the data into the "source"
type Injector interface {
	Inject(data string) (job.Job, error)
}
