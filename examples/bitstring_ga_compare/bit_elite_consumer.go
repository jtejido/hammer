package main

import (
	"fmt"
	"github.com/jtejido/hammer"
)

/**
 * ELITE CONSUMER STRUCT
 * Implements EliteConsumer interface
 **/

type eliteConsumer struct {
	iteration int
}

func (ec *eliteConsumer) OnElite(g *hammer.Individual) {
	gBits := (*g).GetGenome().GetSequence()
	ec.iteration++
	var genomeString string
	for i := 0; i < gBits.Len(); i += 8 {
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
