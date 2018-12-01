package hammer

// GenomeCreator - an interface to a bitset create struct
type GenomeCreator interface {
	Create() Bitset
}

// DefaultBitsetCreate - a null implementation of the IBitsetCreate interface
type DefaultBitsetCreate struct {
}

// Go returns a bitset with no content
func (ngc *DefaultBitsetCreate) Create() Bitset {
	return Bitset{}
}
