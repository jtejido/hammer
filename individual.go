package hammer

type Individual interface {
	GetFitness() int
	SetFitness(int)
	GetGenome() Genome
}

type individual struct {
	fitness int
	genome  Genome
}

// A Individual implementation comprises a population.
func NewIndividual(genome Genome) Individual {
	return &individual{genome: genome}
}
func (i *individual) GetFitness() int {
	return i.fitness
}

func (i *individual) SetFitness(fitness int) {
	i.fitness = fitness
}

func (i *individual) GetGenome() Genome {
	return i.genome
}
