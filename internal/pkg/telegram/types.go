package telegram

import (
	"github.com/shlima/fortune/internal/pkg/bruteforce"
	"github.com/shlima/fortune/internal/pkg/key"
)

type ICli interface {
	SendHeartBeat(heartbit *bruteforce.HeartBit) error
	KeyFound(chain key.Chain) error
}
