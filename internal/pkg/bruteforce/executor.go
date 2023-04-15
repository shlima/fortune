package bruteforce

import (
	"fmt"
	"math"
	"sync"
	"time"

	"github.com/shlima/fortune/internal/pkg/datum"
	"github.com/shlima/fortune/internal/pkg/key"
)

type IopCh = chan int
type FoundFn func(chain key.Chain)

type Executor struct {
	index   datum.Index
	gen     key.IGenerator
	workers int
	sleep   time.Duration
	ch      IopCh
	ops     int64
	t0      time.Time
	prev    int64
	mx      *sync.Mutex
}

func New(index datum.Index, gen key.IGenerator, workers int) *Executor {
	return &Executor{
		index:   index,
		gen:     gen,
		workers: workers,
		ch:      make(IopCh, workers),
		t0:      time.Now(),
		mx:      new(sync.Mutex),
	}
}

// NightMode sets night mode (use less CPU)
func (e *Executor) NightMode(on bool) {
	switch on {
	case true:
		e.sleep = time.Millisecond
	default:
		e.sleep = time.Duration(0)
	}
}

// DataLength returns index items counter
func (e *Executor) DataLength() int {
	return len(e.index)
}

// Get Test the index with the passed address
func (e *Executor) Get(address string) bool {
	return e.index[address]
}

func (e *Executor) Run(fn FoundFn) {
	for i := 0; i < e.workers; i++ {
		go e.run(fn)
	}

	for {
		e.ops += int64(<-e.ch)
	}
}

func (e *Executor) run(foundFn FoundFn) {
	for {
		pair, err := e.gen.Generate()
		if err != nil {
			panic(fmt.Errorf("failed to generate: %w", err))
		}

		if _, ok := e.index[pair.Compressed]; ok {
			foundFn(pair)
		}

		if _, ok := e.index[pair.Uncompressed]; ok {
			foundFn(pair)
		}

		time.Sleep(e.sleep)
		e.ch <- 1
	}
}

func (e *Executor) Heartbeat() *HeartBit {
	e.mx.Lock()
	defer e.mx.Unlock()

	t1 := time.Now()
	tried := e.ops

	out := &HeartBit{
		Tried: tried,
		IOps:  int64(math.Round(float64(tried-e.prev) / t1.Sub(e.t0).Seconds())),
	}

	e.prev = tried
	e.t0 = t1

	return out
}
