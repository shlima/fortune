package bruteforce

import "fmt"

type HeartBit struct {
	Tried int64
	IOps  int64
}

func (h *HeartBit) ToString() string {
	return fmt.Sprintf("tried %d keys (%d ops/sec)", h.Tried, h.IOps)
}
