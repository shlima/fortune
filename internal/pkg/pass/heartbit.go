package pass

import "fmt"

type HeartBit struct {
	Tried    uint64
	IOps     uint64
	Password string
	State    string
}

func (h *HeartBit) ToString() string {
	return fmt.Sprintf(
		"tried %d password (%d ops/sec), current one: <%s>, state: %s",
		h.Tried, h.IOps, h.Password, h.State,
	)
}
