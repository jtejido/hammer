package hammer

// EliteConsumer - an interface to the elite consumer
type EliteConsumer interface {
	OnElite(*Individual)
}

// A nil Elite Consumer
type DefaultEliteConsumer struct {
}

func (nec *DefaultEliteConsumer) OnElite(*Individual) {
}
