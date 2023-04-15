package bruteforce

import (
	"context"
	"fmt"
	"math"
	"sync"
	"sync/atomic"
	"time"

	"github.com/shlima/fortune/internal/pkg/datum"
	"github.com/shlima/fortune/internal/pkg/key"
)

type Executor struct {
	index   datum.Index
	gen     key.IGenerator
	workers int
	sleep   time.Duration
	ch      IopCh
	ops     uint64
	prev    uint64
	wg      *sync.WaitGroup
	t0      time.Time
	mx      *sync.Mutex
	cancel  context.CancelFunc
}

func New(index datum.Index, gen key.IGenerator, workers int) *Executor {
	return &Executor{
		index:   index,
		gen:     gen,
		workers: workers,
		wg:      new(sync.WaitGroup),
		cancel:  func() {},
		t0:      time.Now(),
		mx:      new(sync.Mutex),
	}
}

// SetNightMode sets night mode (use less CPU)
func (e *Executor) SetNightMode(on bool) {
	switch on {
	case true:
		e.sleep = time.Millisecond
	default:
		e.sleep = time.Duration(0)
	}
}

// SetIndex redefined the index
func (e *Executor) SetIndex(index datum.Index) {
	e.mx.Lock()
	defer e.mx.Unlock()

	e.index = index
}

// DataLength returns index items counter
func (e *Executor) DataLength() int {
	return len(e.index)
}

// Get tests the index with the passed address
func (e *Executor) Get(address string) bool {
	return e.index[address]
}

func (e *Executor) Run(fn FoundFn) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	e.cancel = cancel
	e.ch = make(IopCh, e.workers*10_000)
	e.wg = new(sync.WaitGroup)
	for i := 0; i < e.workers; i++ {
		e.wg.Add(1)
		go e.run(ctx, fn)
	}

	for {
		e.SetOpts(uint64(<-e.ch))
	}
}

func (e *Executor) Stop() {
	e.cancel()
	e.wg.Wait()
	close(e.ch)
}

func (e *Executor) run(ctx context.Context, foundFn FoundFn) {
	for {
		select {
		case <-ctx.Done():
			e.wg.Done()
			return
		default:
			e.next(foundFn)
			time.Sleep(e.sleep)
			e.ch <- 1
		}
	}
}

func (e *Executor) next(foundFn FoundFn) {
	pair, err := e.gen.Generate()
	if err != nil {
		panic(fmt.Errorf("failed to generate: %w", err))
	}

	if e.index[pair.Compressed] || e.index[pair.Uncompressed] {
		e.mx.Lock()
		foundFn(pair)
		e.mx.Unlock()
	}
}

func (e *Executor) Heartbeat() *HeartBit {
	e.mx.Lock()
	defer e.mx.Unlock()

	t1 := time.Now()
	tried := e.GetOps()

	out := &HeartBit{
		Tried: tried,
		IOps:  uint64(math.Round(float64(tried-e.prev) / t1.Sub(e.t0).Seconds())),
	}

	e.prev = tried
	e.t0 = t1

	return out
}

func (e *Executor) GetOps() uint64 {
	return atomic.LoadUint64(&e.ops)
}

func (e *Executor) SetOpts(value uint64) {
	atomic.AddUint64(&e.ops, value)
}
