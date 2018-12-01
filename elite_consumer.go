package hammer

// EliteConsumer - an interface to the elite consumer
type EliteConsumer interface {
	OnElite(*Individual)
}

type DefaultEliteConsumer struct {
}

func (nec *DefaultEliteConsumer) OnElite(*Individual) {
}
