package job

type Job interface {
	Revert() error
}
