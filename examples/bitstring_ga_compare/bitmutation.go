package main

import (
	"math/rand"
	. "github.com/jtejido/hammer"
)

/**
 * MUTATION FUNCTION
 * Implements MutationFunc to be used for MutatorFunction interface
 **/

func Mutate(g1 *Individual) Individual {

	g1BitsOrig := (*g1).GetGenome().GetSequence()
	g1Bits := g1BitsOrig.Copy()
	randomBit := rand.Intn(g1Bits.Len())
	g1Bits.Set(randomBit, 1-g1Bits.Get(randomBit).(int))

	return NewIndividual(NewGenome(g1Bits))
}