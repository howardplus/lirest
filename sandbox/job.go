package sandbox

import (
	"time"
)

// Job represents either a delayed job
// or a revertible job
type Job struct {
	Time time.Time
	Id   int
}
