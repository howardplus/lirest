package inject

// Injector implements the Inject method
// it validate and writes the data into the "source"
type Injector interface {
	Inject(data string) error
	Name() string
}
