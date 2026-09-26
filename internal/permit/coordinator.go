package permit

type Evaluator struct {
	generation map[string]int64
	open       map[string]bool
}

func NewEvaluator() *Evaluator {
	return &Evaluator{generation: make(map[string]int64), open: make(map[string]bool)}
}
