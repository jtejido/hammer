package ga

import (
	"fmt"
	. "github.com/jtejido/hammer"
)

// GA Implementation of the Evolution interface, this is almost what you need for a proper GA population generator.
// * Reproduction - The component used that accepts two individuals, do crossover and mutate, then give two new offspring.
// * Selector - The selection component used to pick a random individual.
// * SequenceCreator - used to create the initial population of genome splice.
type GAEvolution struct {
	Selector        Selector
	Reproduction    Reproduction
	SequenceCreator SequenceCreator
}

func NewEvolution(selector Selector, reproduction Reproduction, sequenceCreator SequenceCreator) Evolution {
	return &GAEvolution{
		Selector:        selector,
		Reproduction:    reproduction,
		SequenceCreator: sequenceCreator,
	}
}

func (evolution *GAEvolution) Evolve(population []Individual, totalFitness int) ([]Individual, error) {

	newPopulation, err := evolution.Create(len(population))

	if err != nil {
		return nil, err
	}

	for i := 0; i < len(population); i += 2 {
		g1 := evolution.Selector.Select(population, totalFitness)
		g2 := evolution.Selector.Select(population, totalFitness)

		g3, g4 := evolution.Reproduction.Reproduce(g1, g2)

		newPopulation[i] = g3

		if (i + 1) < len(population) {
			newPopulation[i+1] = g4
		}
	}

	return newPopulation, nil

}

func (evolution *GAEvolution) Create(populationSize int) ([]Individual, error) {
	if evolution.SequenceCreator == nil {
		return nil, fmt.Errorf("Error creating individuals.")
	}

	ret := make([]Individual, populationSize)
	for i := 0; i < populationSize; i++ {
		ret[i] = NewIndividual(NewGenome(evolution.SequenceCreator.NewSequence()))
	}
	return ret, nil
}

func (evolution *GAEvolution) OnElite(elite *Individual) {
	evolution.Reproduction.OnElite(elite)
}
