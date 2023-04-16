package blockchain

type ICli interface {
	Addresses([]string) ([]Address, error)
}
