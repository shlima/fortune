package key

import "github.com/shlima/fortune/internal/pkg/domain"

//go:generate mockgen -source types.go -destination ../../mock/key.go -package mock -mock_names IGenerator=MockKeygenerator

type IGenerator interface {
	Generate() (out domain.KeyChain, err error)
	SetTesting(address string) IGenerator
}
