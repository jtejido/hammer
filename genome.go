package hammer

type Genome interface {
	GetBits() *Bitset
}

type genome struct {
	bitset  Bitset
}

func NewGenome(bitset Bitset) Genome {
	return &genome{bitset: bitset}
}

func (g *genome) GetBits() *Bitset {
	return &g.bitset
}
