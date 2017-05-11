package sandbox

import (
	"github.com/howardplus/lirest/inject"
	"time"
)

// Job represents either a delayed job
// or a revertible job
type Job struct {
	Start    time.Time
	Expire   time.Time
	Id       int
	Injector inject.Injector
	Data     string
}

// Revert implements the Job interface
func (j *Job) Revert() error {
	return nil
}
