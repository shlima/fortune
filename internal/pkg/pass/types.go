package pass

type IGen interface {
	Permutations() uint64
	All() []string
	Next() (out string, err error)
	Heartbeat() *HeartBit
	Opts() GenOpts
}
