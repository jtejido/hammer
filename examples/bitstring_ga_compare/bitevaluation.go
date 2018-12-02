package main

import "github.com/jtejido/hammer"

/**
 * EVALUATION STRUCT
 * Implements Evaluator interface
 **/

type bitStringEvaluator struct {
}

func (e *bitStringEvaluator) OnBeginEvaluation() {
}

func (e *bitStringEvaluator) OnEndEvaluation() {
}

func (e *bitStringEvaluator) Evaluate(g *hammer.Individual) {
	bits := (*g).GetGenome().GetSequence()
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

func (e *bitStringEvaluator) Exit(g *hammer.Individual) bool {
	return (*g).GetFitness() == targetLength
}
