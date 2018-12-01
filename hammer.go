package hammer

import (
	"sync"
	"time"
)

const (
	DefaultPopulationSize = 600
)

// EvolutionAlgorithm - The main simulator.
// * Reproduction - combining evolved genomes
// * EliteConsumer - an optional class that accepts the 'elite' of each population generation
// * Evaluator - an evaluation component used to test each individual's fitness.
// * BitsetCreate - used to create the initial population of genomes
type EvolutionAlgorithm struct {
	Reproduction  			Reproduction
	EliteConsumer 			EliteConsumer
	Evaluator     			Evaluator
	Selector      			Selector
	PopulationCreator 		PopulationCreator
	populationSize          int
	threadCount 			int
	population              []Individual
	totalFitness            int
	individualChannel chan *Individual
	exitFunc                func(*Individual) bool
	waitGroup               *sync.WaitGroup
	parallelSimulations     int
}

func New(populationSize int, population []Individual, numThreads int) EvolutionAlgorithm {

	return EvolutionAlgorithm{
		EliteConsumer: &DefaultEliteConsumer{},
		Reproduction:  &DefaultReproduction{},
		Evaluator:     &DefaultEvaluator{},
		Selector:      &DefaultSelector{},
		PopulationCreator:  &DefaultPopulationCreator{
			&DefaultBitsetCreate{},
		},
		populationSize: populationSize,
		population: 	population,
		threadCount: 	numThreads,
		waitGroup:		new(sync.WaitGroup),
	}

}

func (ga *EvolutionAlgorithm) createPopulation() []Individual {
	return ga.PopulationCreator.Populate(ga.populationSize)
}

func (ga *EvolutionAlgorithm) evaluate() {
	ga.Evaluator.OnBeginEvaluation()
	ga.totalFitness = 0

	ga.individualChannel = make(chan *Individual)

	for i := 0; i < ga.threadCount; i++ {
		go func(individualChannel chan *Individual,
			waitGroup *sync.WaitGroup, evaluator Evaluator) {

			for genome := range individualChannel {
				defer waitGroup.Done()
				evaluator.Evaluate(genome)
			}
		}(ga.individualChannel, ga.waitGroup, ga.Evaluator)
	}

	ga.waitGroup.Add(ga.populationSize)
}

func (ga *EvolutionAlgorithm) onNewIndividualToEvaluate(i *Individual) {
	ga.individualChannel <- i
}

func (ga *EvolutionAlgorithm) syncIndividuals() {
	close(ga.individualChannel)
	ga.waitGroup.Wait()
}

func (ga *EvolutionAlgorithm) getElite() *Individual {
	var ret *Individual
	for i := 0; i < ga.populationSize; i++ {
		if ret == nil || ga.population[i].GetFitness() > (*ret).GetFitness() {
			ret = &ga.population[i]
		}
	}
	return ret
}

func (ga *EvolutionAlgorithm) SimulateUntil(exitFunc func(*Individual) bool) bool {
	ga.exitFunc = exitFunc
	return ga.Simulate()
}

func (ga *EvolutionAlgorithm) shouldExit(elite *Individual) bool {
	if ga.exitFunc == nil {
		return ga.Evaluator.Exit(elite)
	}
	return ga.exitFunc(elite)
}

func (ga *EvolutionAlgorithm) Simulate() bool {

	if ga.populationSize == 0 {
		ga.populationSize = DefaultPopulationSize
	}

	if ga.population == nil {
		ga.population = ga.createPopulation()
	}

	ga.evaluate()
	for i := 0; i < ga.populationSize; i++ {
		ga.onNewIndividualToEvaluate(&ga.population[i])
	}
	ga.syncIndividuals()
	ga.Evaluator.OnEndEvaluation()

	for {
		elite := ga.getElite()
		ga.Reproduction.OnElite(elite)
		ga.EliteConsumer.OnElite(elite)
		if ga.shouldExit(elite) {
			break
		}

		time.Sleep(1 * time.Microsecond)

		ga.evaluate()

		newPopulation := ga.createPopulation()
		for i := 0; i < ga.populationSize; i += 2 {
			g1 := ga.Selector.Select(ga.population, ga.totalFitness)
			g2 := ga.Selector.Select(ga.population, ga.totalFitness)

			g3, g4 := ga.Reproduction.Reproduce(g1, g2)

			newPopulation[i] = g3
			ga.onNewIndividualToEvaluate(&newPopulation[i])

			if (i + 1) < ga.populationSize {
				newPopulation[i+1] = g4
				ga.onNewIndividualToEvaluate(&newPopulation[i+1])
			}
		}
		ga.population = newPopulation
		ga.syncIndividuals()
		ga.Evaluator.OnEndEvaluation()
	}

	return true
}

func (ga *EvolutionAlgorithm) GetPopulation() []Individual {
	return ga.population
}
