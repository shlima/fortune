package brainforce

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/shlima/fortune/internal/mock"
	"github.com/shlima/fortune/internal/pkg/datum"
)

type Setup struct {
	key  *mock.MockKeygen
	pass *mock.MockPassgen
	ctrl *gomock.Controller
	*Force
}

func MustSetup(t *testing.T) *Setup {
	ctrl := gomock.NewController(t)
	pass := mock.NewMockPassgen(ctrl)
	key := mock.NewMockKeygen(ctrl)

	return &Setup{
		ctrl:  ctrl,
		pass:  pass,
		key:   key,
		Force: New(make(datum.Index), key, pass, 2),
	}
}
