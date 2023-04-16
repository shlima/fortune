package brainforce

import (
	"errors"
	"fmt"

	"github.com/shlima/fortune/internal/pkg/datum"
	"github.com/shlima/fortune/internal/pkg/key"
	"github.com/shlima/fortune/internal/pkg/pass"
)

type Force struct {
	index datum.Index
	key   key.IGenerator
	pass  pass.IGen
}

func New(index datum.Index, key key.IGenerator, pass pass.IGen) *Force {
	return &Force{
		index: index,
		key:   key,
		pass:  pass,
	}
}

func (f *Force) PassGen() pass.IGen {
	return f.pass
}

func (f *Force) Generate(onFound func(chain key.KeyChain)) error {
LOOP:
	for {
		password, err := f.pass.Next()
		switch {
		case errors.Is(err, pass.ErrEnd):
			break LOOP
		case err != nil:
			return fmt.Errorf("failed to next: %w", err)
		}

		chain, err := f.key.BrainSHA256([]byte(password))
		if err != nil {
			return fmt.Errorf("failed to key gen <%s>: %w", password, err)
		}

		if f.index[chain.Compressed] || f.index[chain.Uncompressed] {
			onFound(chain)
		}
	}

	return nil
}
