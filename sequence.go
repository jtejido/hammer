package hammer

// Sequence is the slice of data you'll normally work with, be it an []int, []float64, hashmap, etc. is up to you.
type Sequence interface {
	Get(int) interface{}
	Set(int, interface{})
	Copy() Sequence
	Len() int
}

// SequenceCreator - an interface to a sequence create struct implementation
type SequenceCreator interface {
	NewSequence() Sequence
}
