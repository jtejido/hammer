package hammer

// Evolution - an interface to a next population generator implementation
type Evolution interface {
	// evolve and return new individuals from the given population
	Evolve([]Individual, int) ([]Individual, error)
	// Creates an initial set of inividuals
	Create(int) ([]Individual, error)
	OnElite(*Individual)
}
