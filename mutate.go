package hammer

import (
	"math/rand"
)

// MutationFunc - an interface to a mutation func instance
type MutationFunc func(*Individual) Individual

type MutatorFunction interface {
	GetProbability() 	float32
	Mutate(*Individual) Individual
}

// An implementation of MutatorFunction that has a function and a probability
// where function is called with a probability (between 0 to 1)
type Mutator struct {
	Probability        float32
	Func        	   MutationFunc
}

// A helper to initiate a mutator
func NewMutator(probability float32, f MutationFunc) MutatorFunction {
	return &Mutator{
		Probability: probability,
		Func: f,
	}
}

func (m Mutator) GetProbability() float32 {
	return m.Probability
}

func (m Mutator) Mutate(g1 *Individual) Individual {
	return m.Func(g1)
}

// Mutate -
// Accepts an individual and mutates a single bit (flip-bit method) to create a new
// very slightly different individual.
// 000000 to 001000
func Mutate(g1 *Individual) Individual {

	g1BitsOrig := (*g1).GetGenome().GetBits()
	g1Bits := g1BitsOrig.CreateCopy()
	randomBit := rand.Intn(g1Bits.GetSize())
	g1Bits.Set(randomBit, 1-g1Bits.Get(randomBit))

	return NewIndividual(NewGenome(g1Bits))
}