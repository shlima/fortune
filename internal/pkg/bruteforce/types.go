package bruteforce

import "github.com/shlima/fortune/internal/pkg/domain"

type IopCh = chan int
type FoundFn func(chain domain.KeyChain)

type IHeartBit interface {
	ToString() string
}

func EmptyFoundFn(chain domain.KeyChain) {
}
