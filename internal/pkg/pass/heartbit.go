package pass

import "fmt"

type HeartBit struct {
	Tried    uint64
	IOps     uint64
	Password string
	State    string
	DonePct  float64
}

func (h *HeartBit) ToString() string {
	return fmt.Sprintf(
		"tried %d passwords done %.4f%% (%d ops/sec), current one: <%s>, state: %s",
		h.Tried, h.DonePct, h.IOps, h.Password, h.State,
	)
}
