package main

import (
	"github.com/jtejido/hammer"
	"math/rand"
)

/**
 * SEQUENCE CREATOR STRUCT
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
