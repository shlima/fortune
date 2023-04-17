package brainforce

import "github.com/shlima/fortune/internal/pkg/key"

type PassCh chan string
type OnFoundFn = func(chain key.KeyChain)
