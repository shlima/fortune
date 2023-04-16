package pass

import (
	"errors"
	"math"
	"sync"
	"sync/atomic"
	"time"
)

type Gen struct {
	GenOpts
	aMax       int    // alphabet max index
	sMax       int    // state max index
	iterations uint64 // iterations counter
	prev       uint64
	t0         time.Time
	mx         *sync.Mutex
}

type GenOpts struct {
	Alphabet []string
	State    []int // index of each letter in password
	Length   int   // length of the password
}

func New(opts GenOpts) *Gen {
	out := &Gen{
		GenOpts: opts,
		mx:      new(sync.Mutex),
		t0:      time.Now(),
	}

	out.aMax = len(out.Alphabet) - 1
	out.sMax = out.Length - 1

	if len(out.State) == 0 {
		// set state for maximum alphabetical value (max index),
		// as Next function uses reverse looping
		out.State = make([]int, out.Length)
		for i := range out.State {
			out.State[i] = out.aMax
		}
	}

	return out
}

func (g *Gen) Opts() GenOpts {
	return g.GenOpts
}

// Permutations possible number of password variants
func (g *Gen) Permutations() uint64 {
	return uint64(math.Pow(float64(len(g.Alphabet)), float64(g.Length)))
}

// All generates all possible combinations
// (for test usage only)
func (g *Gen) All() []string {
	out := make([]string, 0)
LOOP:
	for {
		pass, err := g.Next()
		switch {
		case err == nil:
			out = append(out, pass)
		case errors.Is(err, ErrEnd):
			break LOOP
		default:
			panic(err)
		}
	}

	return out
}

func (g *Gen) Next() (out string, err error) {
	g.mx.Lock()
	defer g.mx.Unlock()

	if g.addIteration(1) == 1 {
		return MakeOutput(g.Alphabet, g.State), nil
	}

LOOP:
	// going in a reverse order
	for i := g.sMax; i >= 0; i-- {
		g.State[i]--
		switch {
		// end of the first element search
		case i == 0 && g.State[i] == -1:
			break LOOP
		// reset the element t the maximum value
		case g.State[i] == -1:
			g.State[i] = g.aMax
			continue LOOP
		}

		return MakeOutput(g.Alphabet, g.State), nil
	}

	return "", ErrEnd
}

func (g *Gen) Heartbeat() *HeartBit {
	g.mx.Lock()
	defer g.mx.Unlock()

	t1 := time.Now()
	tried := g.GetIterations()

	out := &HeartBit{
		Tried:    tried,
		Password: MakeOutput(g.Alphabet, g.State),
		State:    MarshallState(g.State),
		IOps:     uint64(math.Round(float64(tried-g.prev) / t1.Sub(g.t0).Seconds())),
	}

	g.prev = tried
	g.t0 = t1

	return out
}

func (g *Gen) GetIterations() uint64 {
	return atomic.LoadUint64(&g.iterations) - 1
}

func (g *Gen) addIteration(value uint64) uint64 {
	return atomic.AddUint64(&g.iterations, value)
}
