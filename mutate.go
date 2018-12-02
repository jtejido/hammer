package hammer

// MutationFunc - an interface to a mutation func instance
type MutationFunc func(*Individual) Individual

type MutatorFunction interface {
	GetProbability() float32
	Mutate(*Individual) Individual
}

// An implementation of MutatorFunction that has a function and a probability
// where function is called with a probability (between 0 to 1)
type Mutator struct {
	Probability float32
	Func        MutationFunc
}

// A helper to initiate a mutator
func NewMutator(probability float32, f MutationFunc) MutatorFunction {
	return &Mutator{
		Probability: probability,
		Func:        f,
	}
}

func (m Mutator) GetProbability() float32 {
	return m.Probability
}

func (m Mutator) Mutate(g1 *Individual) Individual {
	return m.Func(g1)
}
