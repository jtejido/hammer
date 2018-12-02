package main

import (
	. "github.com/jtejido/hammer"
)

// Bitset - a simple bitset implementation for a Sequence Interface
// This will be used plainly as a slice of 0s and 1s for our sample
type Bitset struct {
	size int
	bits []int
}

// Create - creates a bitset of length 'size'
func (b *Bitset) Create(size int) {
	b.size = size
	b.bits = make([]int, size)
}

// GetSize returns the size of the bitset
func (b Bitset) Len() int {
	return b.size
}

// Get returns the value in the bitset at index 'index'
// or -1 if the index is out of range
func (b Bitset) Get(index int) interface{} {
	if index < b.size {
		return b.bits[index]
	}
	return -1
}

// GetAll returns an int array of all the bits in the bitset
func (b *Bitset) GetAll() []int {
	return b.bits
}

func (b *Bitset) setImpl(index, value int) {

	b.bits[index] = value
}

// Set assigns value 'value' to the bit at index 'index'
func (b Bitset) Set(index int, value interface{}) {
	if index < b.size {
		b.setImpl(index, value.(int))
	}
}

// SetAll assigns the value 'value' to all the bits in the set
func (b *Bitset) SetAll(value int) {
	for i := 0; i < b.size; i++ {
		b.setImpl(i, value)
	}
}

// CreateCopy returns a bit for bit copy of the bitset
func (b Bitset) Copy() Sequence {
	newBitset := Bitset{}
	newBitset.Create(b.size)
	for i := 0; i < b.size; i++ {
		newBitset.Set(i, b.Get(i))
	}
	return newBitset
}

// Slice returns an out of memory slice of the current bitset
// between bits 'startingBit' and 'startingBit + size'
func (b *Bitset) Slice(startingBit, size int) Bitset {
	ret := Bitset{}
	ret.Create(size)
	ret.bits = b.bits[startingBit : startingBit+size]
	return ret
}
