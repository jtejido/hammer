package hammer

import (
	"sync"
	"time"
)

const (
	DefaultPopulationSize = 600
)

// EvolutionAlgorithm - The main simulator.
// * EliteConsumer - an optional class that accepts the 'elite' of each population generation.
// * Evaluator - an evaluation component used to test each individual's fitness.
// * Evolver - an EA algorithm component used for generating and evolving population.
type EvolutionAlgorithm struct {
	EliteConsumer EliteConsumer
	Evaluator     Evaluator
	Evolver       Evolution

	populationSize    int                    // optional, when population isn't provided, used for creating population.
	threadCount       int                    // used for evaluation concurrency
	population        []Individual           // group of inidividuals
	totalFitness      int                    // optional
	individualChannel chan *Individual       // used for syncing individuals
	exitFunc          func(*Individual) bool // used as termination function per evaluation of an individual
	waitGroup         *sync.WaitGroup        // used for syncing individuals
}

func New(populationSize int, population []Individual, numThreads int) EvolutionAlgorithm {

	return EvolutionAlgorithm{
		EliteConsumer:  &DefaultEliteConsumer{},
		Evaluator:      &DefaultEvaluator{},
		Evolver:        nil,
		populationSize: populationSize,
		population:     population,
		threadCount:    numThreads,
		waitGroup:      new(sync.WaitGroup),
	}

}

// evaluate each individual concurrently
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
		pop, err_p := ga.Evolver.Create(ga.populationSize)

		if err_p != nil {
			return false
		}

		ga.population = pop
	}

	if ga.Evolver == nil {
		return false
	}

	ga.evaluate()
	for i := 0; i < ga.populationSize; i++ {
		ga.onNewIndividualToEvaluate(&ga.population[i])
	}
	ga.syncIndividuals()
	ga.Evaluator.OnEndEvaluation()

	for {
		elite := ga.getElite()
		ga.Evolver.OnElite(elite)
		ga.EliteConsumer.OnElite(elite)
		if ga.shouldExit(elite) {
			break
		}

		time.Sleep(1 * time.Microsecond)

		ga.evaluate()

		newPopulation, err := ga.Evolver.Evolve(ga.population, ga.totalFitness)

		if err != nil {
			return false
		}

		for i := 0; i < ga.populationSize; i++ {
			ga.onNewIndividualToEvaluate(&newPopulation[i])
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
