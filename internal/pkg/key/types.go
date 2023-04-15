package key

type IGenerator interface {
	Generate() (out Chain, err error)
	SetTesting(address string) IGenerator
}
