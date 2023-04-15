package telegram

import (
	"github.com/shlima/fortune/internal/pkg/bruteforce"
	"github.com/shlima/fortune/internal/pkg/key"
)

type NoOp struct {
}

func NewNoOp() *NoOp {
	return &NoOp{}
}

func (n *NoOp) HeartBeat(heartbit *bruteforce.HeartBit) error {
	return nil
}

func (n *NoOp) KeyFound(chain key.Chain) error {
	return nil
}
