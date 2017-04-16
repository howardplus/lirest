package util

type VersionTag struct {
	Major int
	Minor int
}

var Version VersionTag = VersionTag{Major: 0, Minor: 1}
