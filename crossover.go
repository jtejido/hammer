package hammer

// CrossOverFunc - an interface to a crossover func instance
type CrossOverFunc func(*Individual, *Individual) (Individual, Individual)

type CrossOverFunction interface {
	GetProbability() 					float32
	IsUseElite() 						bool
	CrossOver(*Individual, *Individual) (Individual, Individual)
}

// An implementation of CrossOverFunction that has a function and a probability
// where function is called with a probability (between 0 to 1)
type CrossOver struct {
	Probability float32
	Func        CrossOverFunc
	UseElite 	bool
}

// A helper struct to initiate a crossover
func NewCrossOver(probability float32, useElite bool, f CrossOverFunc) CrossOverFunction {
	return &CrossOver{
		Probability: probability,
		Func: f,
		UseElite: useElite,
	}
}

func (m CrossOver) GetProbability() float32 {
	return m.Probability
}

func (m CrossOver) IsUseElite() bool {
	return m.UseElite
}

func (m CrossOver) CrossOver(g1, g2 *Individual) (Individual, Individual) {
	return m.Func(g1, g2)
}