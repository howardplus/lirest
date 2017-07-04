package describe

// DescDefn is the description definition map
// each type contains a slice of descriptions
type DescDefn struct {
	DescriptionMap map[DescType][]Description
}
