package ga

import (
	"math/rand"
	. "github.com/jtejido/hammer"
)

// GA Implementation of the Reproduction interface, this is almost what you need for a proper GA reproduction.
type GAReproduction struct {
	crossOverFuncs 	[]CrossOverFunction
	mutatorFuncs 	[]MutatorFunction
	elite       	*Individual
}

// NewReproduction returns an implementation of a Reproduction with several CrossOver and Mutator Funcs
func NewReproduction(crossOverFuncs []CrossOverFunction, mutatorFuncs []MutatorFunction) Reproduction {
	return &GAReproduction{
		crossOverFuncs: crossOverFuncs,
		mutatorFuncs: mutatorFuncs,
	}
}

// Reproduce cycles through, and applies, the set crossover and mutation funcs
func (r *GAReproduction) Reproduce(i1, i2 Individual) (Individual, Individual) {

	newInd1 := NewIndividual(NewGenome(i1.GetGenome().GetSequence()))
	newInd2 := NewIndividual(NewGenome(i2.GetGenome().GetSequence()))

	// crossover
	for _, f := range r.crossOverFuncs {
		if rand.Float32() < f.GetProbability() {
			if f.IsUseElite() {
				newInd1, newInd2 = f.CrossOver(&newInd1, r.elite)
			} else {
				newInd1, newInd2 = f.CrossOver(&newInd1, &newInd2)
			}
		}
	}

	// mutate
	for _, f := range r.mutatorFuncs {
		if rand.Float32() < f.GetProbability() {
			newInd1 = f.Mutate(&newInd1)
			newInd2 = f.Mutate(&newInd2)
		}
	}

	return newInd1, newInd2
}

func (r *GAReproduction) OnElite(elite *Individual) {
	r.elite = elite
}