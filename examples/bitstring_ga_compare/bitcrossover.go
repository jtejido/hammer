package main

import (
	. "github.com/jtejido/hammer"
	"math/rand"
)

/**
 * CROSSOVER FUNCTIONS
 * Implements CrossOverFunc to be used for CrossOverFunction Interface
 **/

// OnePointCrossover -
// Accepts 2 genomes and combines them to create 2 new genomes using one point crossover
// 000000 and 111111 to 000111 and 111000
func OnePointCrossover(g1, g2 *Individual) (Individual, Individual) {

	g1Bits, g2Bits := (*g1).GetGenome().GetSequence(), (*g2).GetGenome().GetSequence()

	b1, b2 := Bitset{}, Bitset{}
	g1Size := g1Bits.Len()
	g2Size := g2Bits.Len()
	b1.Create(g1Size)
	b2.Create(g2Size)

	maxSize := max(g1Size, g2Size)
	minSize := min(g1Size, g2Size)
	randIndex := rand.Intn(minSize-1) + 1

	for i := 0; i < randIndex; i++ {
		b1.Set(i, g1Bits.Get(i))
		b2.Set(i, g2Bits.Get(i))
	}

	for i := randIndex; i < minSize; i++ {
		b2.Set(i, g1Bits.Get(i))
		b1.Set(i, g2Bits.Get(i))
	}

	if g1Size > g2Size {
		for i := minSize; i < maxSize; i++ {
			b2.Set(i, g1Bits.Get(i))
		}
	} else {
		for i := minSize; i < maxSize; i++ {
			b1.Set(i, g2Bits.Get(i))
		}
	}

	return NewIndividual(NewGenome(b1)), NewIndividual(NewGenome(b2))
}

// TwoPointCrossover -
// Accepts 2 genomes and combines them to create 2 new genomes using two point crossover
// 000000 and 111111 to 001100 and 110011
func TwoPointCrossover(g1, g2 *Individual) (Individual, Individual) {

	g1Bits, g2Bits := (*g1).GetGenome().GetSequence(), (*g2).GetGenome().GetSequence()

	b1, b2 := Bitset{}, Bitset{}
	g1Size := g1Bits.Len()
	g2Size := g2Bits.Len()
	b1.Create(g1Size)
	b2.Create(g2Size)

	maxSize := max(g1Size, g2Size)
	minSize := min(g1Size, g2Size)
	randIndex1 := rand.Intn(minSize-1) + 1
	randIndex2 := randIndex1

	for randIndex1 == randIndex2 {
		randIndex2 = rand.Intn(minSize-1) + 1
	}

	if randIndex1 > randIndex2 {
		randIndex1, randIndex2 = randIndex2, randIndex1
	}

	for i := 0; i < randIndex1; i++ {
		b1.Set(i, g1Bits.Get(i))
		b2.Set(i, g2Bits.Get(i))
	}

	for i := randIndex1; i < randIndex2; i++ {
		b2.Set(i, g1Bits.Get(i))
		b1.Set(i, g2Bits.Get(i))
	}

	for i := randIndex2; i < minSize; i++ {
		b1.Set(i, g1Bits.Get(i))
		b2.Set(i, g2Bits.Get(i))
	}

	if g1Size > g2Size {
		for i := minSize; i < maxSize; i++ {
			b2.Set(i, g1Bits.Get(i))
		}
	} else {
		for i := minSize; i < maxSize; i++ {
			b1.Set(i, g2Bits.Get(i))
		}
	}
	return NewIndividual(NewGenome(b1)), NewIndividual(NewGenome(b2))
}

// UniformCrossover -
// Accepts 2 genomes and combines them to create 2 new genomes using uniform crossover
// 000000 and 111111 to 101010 and 010101
func UniformCrossover(g1, g2 *Individual) (Individual, Individual) {

	g1Bits, g2Bits := (*g1).GetGenome().GetSequence(), (*g2).GetGenome().GetSequence()

	b1, b2 := Bitset{}, Bitset{}
	g1Size := g1Bits.Len()
	g2Size := g2Bits.Len()
	b1.Create(g1Size)
	b2.Create(g2Size)

	maxSize := max(g1Size, g2Size)
	minSize := min(g1Size, g2Size)

	for i := 0; i < minSize; i++ {
		if rand.Float32() > 0.5 {
			b1.Set(i, g1Bits.Get(i))
			b2.Set(i, g2Bits.Get(i))
		} else {
			b2.Set(i, g1Bits.Get(i))
			b1.Set(i, g2Bits.Get(i))
		}
	}

	if g1Size > g2Size {
		for i := minSize; i < maxSize; i++ {
			b2.Set(i, g1Bits.Get(i))
		}
	} else {
		for i := minSize; i < maxSize; i++ {
			b1.Set(i, g2Bits.Get(i))
		}
	}

	return NewIndividual(NewGenome(b1)), NewIndividual(NewGenome(b2))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
