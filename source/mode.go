package source

const (
	// ModeReadOnly is read only
	ModeReadOnly = iota
	// ModeReadWrite is read/write
	ModeReadWrite
)

// AccessMode is either read or write
type AccessMode int
