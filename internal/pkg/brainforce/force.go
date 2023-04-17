package brainforce

import (
	"errors"
	"fmt"
	"sync"

	"github.com/samber/lo"
	"github.com/shlima/fortune/internal/pkg/datum"
	"github.com/shlima/fortune/internal/pkg/key"
	"github.com/shlima/fortune/internal/pkg/pass"
)

type Force struct {
	index    datum.Index
	key      key.IGen
	pass     pass.IGen
	workers  int
	wg       *sync.WaitGroup
	mx       *sync.Mutex
	channels []PassCh
	onFound  OnFoundFn
}

func New(index datum.Index, key key.IGen, pass pass.IGen, workers int) *Force {
	return &Force{
		key:     key,
		pass:    pass,
		index:   index,
		mx:      new(sync.Mutex),
		wg:      new(sync.WaitGroup),
		workers: workers,
		channels: lo.Map(make([]PassCh, workers), func(item PassCh, index int) PassCh {
			return make(PassCh, 2)
		}),
	}
}

func (f *Force) SetIndex(index datum.Index) {
	f.index = index
}

func (f *Force) PassGen() pass.IGen {
	return f.pass
}

func (f *Force) DataLength() int {
	return len(f.index)
}

// Get tests the index with the passed address
func (f *Force) Get(address string) bool {
	return f.index[address]
}

func (f *Force) Generate(onFound OnFoundFn) error {
	f.onFound = onFound
	f.asyncWatch()

LOOP:
	for i := 0; i <= len(f.channels); i++ {
		if f.workers == i {
			goto LOOP
		}

		password, err := f.pass.Next()
		switch {
		case errors.Is(err, pass.ErrEnd):
			break LOOP
		case err != nil:
			return fmt.Errorf("failed to next: %w", err)
		}

		f.channels[i] <- password
	}

	f.stop()
	return nil
}

func (f *Force) asyncWatch() {
	for i := range f.channels {
		f.wg.Add(1)
		go f.watch(f.channels[i])
	}
}

func (f *Force) stop() {
	for i := range f.channels {
		close(f.channels[i])
	}

	f.wg.Wait()
}

func (f *Force) watch(ch PassCh) {
	for password := range ch {
		chain, err := f.key.BrainSHA256([]byte(password))
		if err != nil {
			panic(fmt.Errorf("failed to key gen <%s>: %w", password, err))
		}

		if f.index[chain.Compressed] || f.index[chain.Uncompressed] {
			f.mx.Lock()
			f.onFound(chain)
			f.mx.Unlock()
		}
	}

	f.wg.Done()
}
