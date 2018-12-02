package main

import (
	"fmt"
	"github.com/jtejido/hammer"
	"github.com/jtejido/hammer/ga"
	"runtime"
	"runtime/debug"
	"time"
)

const (
	populationSize = 600
)

func printMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

var (
	// our goal
	targetString = "Hello Earthlings! We come in pieces.."
	targetLength = len(targetString) * 8
)

/**
 * THIS IS A SAMPLE GA PROGRAM THAT USES BITSET AS GENOME SEQUENCE TO SLOWLY CONVERGE TOWARDS OUR STRING GOAL.
 * THE FILES HAVE BEEN SEPARATED FOR A CLEANER STRUCTURE. PLEASE CHECK OUT HOW EACH INJECTIBLES ARE CREATED.
 **/

func main() {
	debug.SetGCPercent(100)
	numThreads := 4
	runtime.GOMAXPROCS(numThreads)

	h := hammer.New(populationSize, nil, numThreads)

	// inject evaluator
	h.Evaluator = &bitStringEvaluator{}

	// inject elite consumer
	h.EliteConsumer = &eliteConsumer{}

	// The main EA algorithm we wish to use.
	// inject evolver requirements (crossover, mutation and selection, and a sequence creator)
	h.Evolver = ga.NewEvolution(
		hammer.NewSelector(
			[]hammer.SelectorFunction{
				{Probability: 1.0, Func: hammer.Roulette},
			},
		),
		ga.NewReproduction(
			[]hammer.CrossOverFunction{
				hammer.NewCrossOver(1.0, true, UniformCrossover),
			},
			[]hammer.MutatorFunction{
				hammer.NewMutator(1.0, Mutate),
			},
		),
		&sequenceCreator{},
	)

	startTime := time.Now()

	// start
	h.Simulate()

	fmt.Println(time.Since(startTime))
	printMemUsage()
}
