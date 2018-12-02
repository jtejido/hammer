package hammer

// Reproduction - an interface to a reproduction instance
type Reproduction interface {
	Reproduce(Individual, Individual) (Individual, Individual)
	OnElite(*Individual)
}

// DefaultReproduction - default implementation of the Reproduction interface
type DefaultReproduction struct {
}

// Reproduce - default implementation of the Reproduction Reproduce func
func (dm *DefaultReproduction) Reproduce(a, b Individual) (Individual, Individual) {
	return NewIndividual(NewGenome(a.GetGenome().GetSequence())), NewIndividual(NewGenome(b.GetGenome().GetSequence()))
}

// OnElite - default implementation of the Reproduction OnElite func
func (dm *DefaultReproduction) OnElite(a *Individual) {
}