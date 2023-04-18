package pass

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMarshallState(t *testing.T) {
	t.Parallel()

	t.Run("it works", func(t *testing.T) {
		got := MarshallState([]int{})
		require.Equal(t, "", got)

		got = MarshallState([]int{1, 2})
		require.Equal(t, "1,2", got)
	})
}

func TestUnmarshalState(t *testing.T) {
	t.Parallel()

	t.Run("it works", func(t *testing.T) {
		got, err := UnmarshalState("")
		require.NoError(t, err)
		require.Equal(t, []int{}, got)

		got, err = UnmarshalState("1,2")
		require.NoError(t, err)
		require.Equal(t, []int{1, 2}, got)
	})
}

func TestMarshallAlphabet(t *testing.T) {
	t.Parallel()

	t.Run("it works", func(t *testing.T) {
		got := MarshallAlphabet([]string{"a", "b"})
		require.Equal(t, "ab", got)
	})
}

func TestDonePct(t *testing.T) {
	t.Parallel()

	t.Run("when zero", func(t *testing.T) {
		require.EqualValues(t, 0, DonePct(nil, 0))
	})

	t.Run("case#1", func(t *testing.T) {
		t.Parallel()

		alphabetSize := 20
		state0 := []int{19, 19, 19}
		state50 := []int{10, 0, 0}
		state100 := []int{0, 0, 0}

		got := DonePct(state0, alphabetSize)
		require.EqualValues(t, 0, got, "1/8000")

		got = DonePct(state50, alphabetSize)
		require.EqualValues(t, 50.0, got, "4000/8000")

		got = DonePct(state100, alphabetSize)
		require.EqualValues(t, 100, got, "0/8000")
	})

	t.Run("case#2", func(t *testing.T) {
		t.Parallel()

		// 50% alphabet=20,iterations=4000,permutations=8000
		state := []int{10, 0, 0}
		got := DonePct(state, 20)
		require.EqualValues(t, 50.0, got, "4000/8000")
	})

	t.Run("case#3", func(t *testing.T) {
		t.Parallel()

		// [16 16 10] = 1270, alphabet=20, permutations: 8000, state: 19,19,19
		state := []int{16, 16, 10}
		got := DonePct(state, 20)
		require.EqualValues(t, 15.875, got, "1270/8000")
	})

	t.Run("case#4", func(t *testing.T) {
		t.Parallel()

		// [16 16 10] = 10, alphabet=20, permutations: 8000, state: 19,19,19
		state := []int{19, 19, 10}
		got := DonePct(state, 20)
		require.EqualValues(t, 0.125, got, "10/8000")
	})

	t.Run("case#5", func(t *testing.T) {
		t.Parallel()

		// iterations=24062 alphabet=20, permutations: 3200000, state: 19,19,19,19,19
		state := []int{19, 16, 19, 16, 18}
		got := DonePct(state, 20)
		require.EqualValues(t, 0.7519375, got, "24062/3200000")
	})

	t.Run("when 0%", func(t *testing.T) {
		t.Parallel()

		// alphabet=20, permutations: 3200000, state: 19,19,19,19,19
		state := []int{19, 19, 19, 19, 19}
		got := DonePct(state, 20)
		require.EqualValues(t, 0, got)
	})

	t.Run("when 100%", func(t *testing.T) {
		t.Parallel()

		// alphabet=20, permutations: 3200000, state: 19,19,19,19,19
		state := []int{0, 0, 0, 0, 0}
		got := DonePct(state, 20)
		require.EqualValues(t, 100, got)
	})
}
