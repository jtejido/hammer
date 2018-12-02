package main

import (
	"math/rand"
	"github.com/jtejido/hammer"
)

/**
 * SEQUENCE STRUCT
 * Implements SequenceCreator interface
 **/

type sequenceCreator struct {
}

func (bc *sequenceCreator) NewSequence() hammer.Sequence {
	b := Bitset{}
	b.Create(targetLength)
	for i := 0; i < targetLength; i++ {
		b.Set(i, rand.Intn(2))
	}
	return b
}