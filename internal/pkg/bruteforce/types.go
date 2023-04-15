package bruteforce

import "github.com/shlima/fortune/internal/pkg/domain"

type CloseCh = chan bool
type FoundFn func(chain domain.KeyChain)

type IHeartBit interface {
	ToString() string
}

func EmptyFoundFn(chain domain.KeyChain) {
}
