package hammer

// PopulationCreator
type PopulationCreator interface {
	Populate(int) []Individual
}

type DefaultPopulationCreator struct {
	GenomeCreatorFunc 	GenomeCreator
}

func NewPopulation(f GenomeCreator) PopulationCreator {
	return &DefaultPopulationCreator{
		GenomeCreatorFunc: f,
	}
}

// Populate returns a group of individuals
func (p *DefaultPopulationCreator) Populate(populationSize int) []Individual {
	ret := make([]Individual, populationSize)
	for i := 0; i < populationSize; i++ {
		ret[i] = NewIndividual(NewGenome(p.GenomeCreatorFunc.Create()))
	}
	return ret
}
