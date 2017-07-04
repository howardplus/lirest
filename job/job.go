package job

// Job interface that supports a Revert() function
type Job interface {
	Revert() error
}
