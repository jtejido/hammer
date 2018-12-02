package hammer

type Genome interface {
	GetSequence() Sequence
}

type genome struct {
	sequence Sequence
}

func NewGenome(sequence Sequence) Genome {
	return &genome{sequence: sequence}
}

func (g genome) GetSequence() Sequence {
	return g.sequence
}
