package key

//go:generate mockgen -source types.go -destination ../../mock/key.go -package mock -mock_names IGenerator=MockKeygenerator

type IGenerator interface {
	Generate() (out KeyChain, err error)
	SetTesting(address string) IGenerator
	BrainSHA256(password []byte) (out KeyChain, err error)
}
