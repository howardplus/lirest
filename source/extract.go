package source

// an extractor returns a generic data based
// on the converter.
// An object that implements the Extract() interface needs
// to know where to get the data, which then feeds to the
// converter.
type Extractor interface {
	Extract(conv Converter) (interface{}, error)
}
