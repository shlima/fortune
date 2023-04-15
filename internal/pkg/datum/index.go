package datum

import (
	"math/rand"
	"time"
)

type Index map[string]bool

// Random returns random address from the index
func (i Index) Random() (address string, ix int) {
	if len(i) == 0 {
		return
	}

	number := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(len(i))

LOOP:
	for key, _ := range i {
		switch number {
		case ix:
			address = key
			break LOOP
		default:
			ix++
		}
	}

	return
}
