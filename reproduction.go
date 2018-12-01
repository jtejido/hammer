package hammer

import (
	"math/rand"
)

// Reproduction - an interface to a reproduction instance
type Reproduction interface {
	Reproduce(*Individual, *Individual) (Individual, Individual)
	OnElite(*Individual)
}

// DefaultReproduction - default implementation of the Reproduction interface
type DefaultReproduction struct {
}

// Reproduce - default implementation of the Reproduction go func
func (dm *DefaultReproduction) Reproduce(a, b *Individual) (Individual, Individual) {
	return NewIndividual(NewGenome(*(*a).GetGenome().GetBits())), NewIndividual(NewGenome(*(*b).GetGenome().GetBits()))
}

// OnElite - default implementation of the Reproduction OnElite func
func (dm *DefaultReproduction) OnElite(a *Individual) {
}

type reproduction struct {
	crossOverFuncs 	[]CrossOverFunction
	mutatorFuncs 	[]MutatorFunction
	elite       	*Individual
}

// NewReproduction returns an implementation of a Reproduction with several CrossOver and Mutator Funcs
func NewReproduction(crossOverFuncs []CrossOverFunction, mutatorFuncs []MutatorFunction) Reproduction {
	return &reproduction{
		crossOverFuncs: crossOverFuncs,
		mutatorFuncs: mutatorFuncs,
	}
}

// Reproduce cycles through, and applies, the set crossover and mutation funcs
func (r *reproduction) Reproduce(i1, i2 *Individual) (Individual, Individual) {

	newInd1 := NewIndividual(NewGenome(*(*i1).GetGenome().GetBits()))
	newInd2 := NewIndividual(NewGenome(*(*i2).GetGenome().GetBits()))
	
	// crossover
	for _, config := range r.crossOverFuncs {
		if rand.Float32() < config.GetProbability() {
			if config.IsUseElite() {
				newInd1, newInd2 = config.CrossOver(&newInd1, r.elite)
			} else {
				newInd1, newInd2 = config.CrossOver(&newInd1, &newInd2)
			}
		}
	}

	// mutate
	for _, config := range r.mutatorFuncs {
		if rand.Float32() < config.GetProbability() {
			newInd1 = config.Mutate(&newInd1)
			newInd2 = config.Mutate(&newInd2)
		}
	}

	return newInd1, newInd2
}

func (r *reproduction) OnElite(elite *Individual) {
	r.elite = elite
}