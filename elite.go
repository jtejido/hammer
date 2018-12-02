package hammer

// EliteConsumer - an interface to the elite consumer
type EliteConsumer interface {
	OnElite(*Individual)
}

// A Default Elite Consumer that does nothing.
type DefaultEliteConsumer struct {
}

func (nec *DefaultEliteConsumer) OnElite(*Individual) {
}
