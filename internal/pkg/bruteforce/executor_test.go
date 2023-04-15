package bruteforce

import (
	"testing"
	"time"

	"github.com/shlima/fortune/internal/pkg/datum"
	"github.com/shlima/fortune/internal/pkg/domain"
	"github.com/stretchr/testify/require"
)

func TestExecutor_SetNightMode(t *testing.T) {
	t.Parallel()

	t.Run("it works", func(t *testing.T) {
		setup := MustSetup(t, nil)
		defer setup.ctrl.Finish()

		require.EqualValues(t, 0, setup.sleep)
		setup.SetNightMode(true)
		require.EqualValues(t, time.Millisecond, setup.sleep)
		setup.SetNightMode(false)
		require.EqualValues(t, 0, setup.sleep)
	})
}

func TestExecutor_DataLength(t *testing.T) {
	t.Parallel()

	t.Run("it works", func(t *testing.T) {
		setup := MustSetup(t, datum.Index{"foo": true})
		defer setup.ctrl.Finish()

		require.Equal(t, 1, setup.DataLength())
	})
}

func TestExecutor_RunAsync(t *testing.T) {
	t.Parallel()

	t.Run("when nothing found", func(t *testing.T) {
		t.Parallel()

		setup := MustSetup(t, make(datum.Index))
		defer setup.ctrl.Finish()

		setup.gen.EXPECT().
			Generate().
			Return(domain.KeyChain{}, nil).
			AnyTimes()

		got := make([]domain.KeyChain, 0)
		setup.RunAsync(func(chain domain.KeyChain) {
			got = append(got, chain)
		})

		MustSleep(t)
		setup.Stop()
		require.Empty(t, got)
	})

	t.Run("when Compressed found", func(t *testing.T) {
		t.Parallel()
		setup := MustSetup(t, nil)
		defer setup.ctrl.Finish()

		chain := domain.KeyChain{
			Compressed: "Compressed",
		}

		setup.SetIndex(datum.Index{chain.Compressed: true})

		setup.gen.EXPECT().
			Generate().
			Return(chain, nil).
			AnyTimes()

		got := make([]domain.KeyChain, 0)
		setup.RunAsync(func(chain domain.KeyChain) {
			got = append(got, chain)
		})

		MustSleep(t)
		setup.Stop()
		require.NotEmpty(t, got)
		require.Equal(t, chain, got[0])
	})

	t.Run("when Uncompressed found", func(t *testing.T) {
		t.Parallel()
		setup := MustSetup(t, nil)
		defer setup.ctrl.Finish()

		chain := domain.KeyChain{
			Uncompressed: "Uncompressed",
		}

		setup.SetIndex(datum.Index{chain.Uncompressed: true})

		setup.gen.EXPECT().
			Generate().
			Return(chain, nil).
			AnyTimes()

		got := make([]domain.KeyChain, 0)
		setup.RunAsync(func(chain domain.KeyChain) {
			got = append(got, chain)
		})

		MustSleep(t)
		setup.Stop()
		require.NotEmpty(t, got)
		require.Equal(t, chain, got[0])
	})
}

func TestExecutor_Heartbeat(t *testing.T) {
	t.Parallel()

	t.Run("it works", func(t *testing.T) {
		setup := MustSetup(t, nil)
		defer setup.ctrl.Finish()

		setup.gen.EXPECT().
			Generate().
			Return(domain.KeyChain{}, nil).
			AnyTimes()

		setup.RunAsync(EmptyFoundFn)
		MustSleep(t)
		setup.Stop()

		got := setup.Heartbeat()
		require.NotEmpty(t, got.IOps)
		require.NotEmpty(t, got.Tried)
	})
}

func TestExecutor_Run(t *testing.T) {
	t.Run("it works", func(t *testing.T) {
		setup := MustSetup(t, nil)
		defer setup.ctrl.Finish()

		setup.SetWorkers(0)
		setup.Run(EmptyFoundFn)
	})
}
