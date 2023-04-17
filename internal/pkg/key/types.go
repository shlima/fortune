package key

//go:generate mockgen -source types.go -destination ../../mock/key.go -package mock -mock_names IGen=MockKeygen

type IGen interface {
	Generate() (out KeyChain, err error)
	SetTesting(address string) IGen
	BrainSHA256(password []byte) (out KeyChain, err error)
}
