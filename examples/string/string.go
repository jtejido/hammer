package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"runtime/debug"
	"time"
	"github.com/jtejido/hammer"
)

type testEvaluator struct {
}

func (e *testEvaluator) OnBeginEvaluation() {
}

func (e *testEvaluator) OnEndEvaluation() {
}

func (e *testEvaluator) Evaluate(g *hammer.Individual) {
	bits := (*g).GetGenome().GetBits()
	for i, character := range targetString {
		for j := 0; j < 8; j++ {
			targetBit := character & (1 << uint(j))
			bit := bits.Get((i * 8) + j)
			if targetBit != 0 && bit == 1 {
				(*g).SetFitness((*g).GetFitness() + 1)
			} else if targetBit == 0 && bit == 0 {
				(*g).SetFitness((*g).GetFitness() + 1)
			}
		}
	}

}

func (e *testEvaluator) Exit(g *hammer.Individual) bool {
	return (*g).GetFitness() == targetLength
}

type myGenomeCreator struct {
}

func (bc *myGenomeCreator) Create() hammer.Bitset {
	b := hammer.Bitset{}
	b.Create(targetLength)
	for i := 0; i < targetLength; i++ {
		b.Set(i, rand.Intn(2))
	}
	return b
}

type eliteConsumer struct {
	iteration int
}

func (ec *eliteConsumer) OnElite(g *hammer.Individual) {
	gBits := (*g).GetGenome().GetBits()
	ec.iteration++
	var genomeString string
	for i := 0; i < gBits.GetSize(); i += 8 {
		c := int(0)
		for j := 0; j < 8; j++ {
			bit := gBits.Get(i + j)
			if bit != 0 {
				c |= 1 << uint(j)
			}
		}
		genomeString += string(c)
	}

	fmt.Println(ec.iteration, "\t", genomeString, "\t", (*g).GetFitness())
}

const (
	populationSize = 600
)

var (
	targetString = "Hello Earthlings! We come in pieces.."
	targetLength = len(targetString) * 8
)

func main() {
	debug.SetGCPercent(800)
	numThreads := 4
	runtime.GOMAXPROCS(numThreads)

	h := hammer.New(populationSize, nil, numThreads)

	// inject evaluator
	h.Evaluator = &testEvaluator{}

	// inject population creator with gene creator
	h.PopulationCreator = hammer.NewPopulation(&myGenomeCreator{})

	// inject elite consumer
	h.EliteConsumer = &eliteConsumer{}

	// inject reproduction functions
	h.Reproduction = hammer.NewReproduction(
		[]hammer.CrossOverFunction{
			hammer.NewCrossOver(1.0, true, hammer.UniformCrossover),
		},
		[]hammer.MutatorFunction{
			hammer.NewMutator(1.0, hammer.Mutate),
		},
	)

	// inject selector
	h.Selector = hammer.NewSelector(
		[]hammer.SelectorFunctionProbability{
			{P: 1.0, F: hammer.Roulette},
		},
	)


	startTime := time.Now()

	// start
	h.Simulate()

	fmt.Println(time.Since(startTime))
	hammer.PrintMemUsage()
}


