package bruteforce

import (
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/shlima/fortune/internal/mock"
	"github.com/shlima/fortune/internal/pkg/datum"
)

type Setup struct {
	ctrl *gomock.Controller
	gen  *mock.MockKeygenerator
	*Executor
}

func MustSetup(t *testing.T, index datum.Index) *Setup {
	ctrl := gomock.NewController(t)
	gen := mock.NewMockKeygenerator(ctrl)

	return &Setup{
		ctrl:     ctrl,
		gen:      gen,
		Executor: New(index, gen, 2),
	}
}

func MustSleep(t *testing.T) {
	time.Sleep(time.Millisecond * 100)
}
