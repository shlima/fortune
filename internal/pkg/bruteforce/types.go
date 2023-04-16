package bruteforce

import (
	"github.com/shlima/fortune/internal/pkg/key"
)

type CloseCh = chan bool
type FoundFn func(chain key.KeyChain)

type IHeartBit interface {
	ToString() string
}

func EmptyFoundFn(chain key.KeyChain) {
}
