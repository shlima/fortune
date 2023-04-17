package pass

//go:generate mockgen -source types.go -destination ../../mock/pass.go -package mock -mock_names IGen=MockPassgen
type IGen interface {
	Permutations() uint64
	All() []string
	Next() (out string, err error)
	Heartbeat() *HeartBit
	Opts() GenOpts
}
