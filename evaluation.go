package hammer

// Evaluation - an Evaluation interface used to test fitness
type Evaluator interface {
	OnBeginEvaluation()
	Evaluate(*Individual)
	OnEndEvaluation()
	Exit(*Individual) bool
}

type DefaultEvaluator struct {
}

func (ns *DefaultEvaluator) Evaluate(*Individual) {
}

func (ns *DefaultEvaluator) OnBeginEvaluation() {
}

func (ns *DefaultEvaluator) OnEndEvaluation() {
}

func (ns *DefaultEvaluator) Exit(*Individual) bool {
	return false
}
