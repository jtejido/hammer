package hammer

import (
	"math/rand"
)

// Selector - a selector interface used to pick 2 genomes to mate
type Selector interface {
	Select([]Individual, int) Individual
}

// DefaultSelector - a default implementation of the Selector interface
type DefaultSelector struct {
}

// Select - a null implementation of Selector's 'go'
func (s *DefaultSelector) Select(individuals []Individual, totalFitness int) Individual {
	return individuals[0]
}

// SelectorFunction -
// Contains a selector function and a probability
// where selector function 'F' is called with probability 'P'
// where 'P' is a value between 0 and 1
// 0 = never called, 1 = called every time we need a new genome to mate
type SelectorFunction struct {
	Probability float32
	Func func([]Individual, int) Individual
}

type selector struct {
	selectorConfig []SelectorFunction
}

// NewSelector returns an instance of an Selector with several SelectorFunctionProbabiities
func NewSelector(selectorConfig []SelectorFunction) Selector {
	return &selector{
		selectorConfig: selectorConfig,
	}
}

// Select - cycles through the selector function probabilities until one returns a genome
func (s *selector) Select(genomeArray []Individual, totalFitness int) Individual {
	for {
		for _, config := range s.selectorConfig {
			if rand.Float32() < config.Probability {
				return config.Func(genomeArray, totalFitness)
			}
		}
	}
}

// Roulette is a selection function that selects a genome where genomes that have a higher fitness are more likely to be picked
func Roulette(individuals []Individual, totalFitness int) Individual {

	if len(individuals) == 0 {
		panic("population contains no individuals.")
	}

	if totalFitness == 0 {
		randomIndex := rand.Intn(len(individuals))
		return individuals[randomIndex]
	}

	randomFitness := rand.Intn(totalFitness)
	for i := range individuals {
		randomFitness -= individuals[i].GetFitness()
		if randomFitness <= 0 {
			return individuals[i]
		}
	}

	panic("total fitness is too large")
}