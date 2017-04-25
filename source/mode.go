package source

const (
	ModeReadOnly  = iota // read-only value
	ModeReadWrite        // writable value
)

// AccessMode: either read or write
type AccessMode int
