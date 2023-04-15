package bruteforce

import "fmt"

type HeartBit struct {
	Tried uint64
	IOps  uint64
}

func (h *HeartBit) ToString() string {
	return fmt.Sprintf("tried %d keys (%d ops/sec)", h.Tried, h.IOps)
}
