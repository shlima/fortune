package telegram

type NoOp struct {
}

func NewNoOp() *NoOp {
	return &NoOp{}
}

func (n *NoOp) HeartBeat(message string) error {
	return nil
}

func (n *NoOp) KeyFound(message string) error {
	return nil
}
