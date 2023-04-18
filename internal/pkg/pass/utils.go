package pass

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func MakeOutput(alphabet []string, state []int) string {
	out := make([]string, len(state))
	for i, num := range state {
		out[i] = alphabet[num]
	}

	return strings.Join(out, "")
}

func MarshallAlphabet(in []string) string {
	return strings.Join(in, "")
}

func MarshallState(in []int) string {
	return strings.Trim(strings.Replace(fmt.Sprint(in), " ", ",", -1), "[]")
}

func UnmarshalState(in string) (out []int, err error) {
	if in == "" {
		return make([]int, 0), nil
	}

	got := strings.Split(in, ",")
	out = make([]int, len(got))
	for i := range got {
		out[i], err = strconv.Atoi(got[i])
		if err != nil {
			return out, fmt.Errorf("failed to parse: %w", err)
		}
	}

	return out, nil
}

func Permutations(alphabetSize, passwordLength int) uint64 {
	return uint64(math.Pow(float64(alphabetSize), float64(passwordLength)))
}

// DonePct returns % of done work
//
//	[16 16 10] = 1270, alphabet=20, permutations: 8000, state: 19,19,19
//	[20 20 20]
//	 400 20 1  iterations per char in state position
//	[03 03 09]
//	 400*3 + 20*3 + 9 = 1269 (iterations)
//
//	alphabet: 1,2,3,4
//	length: 3
//	[4,4,4]
//	[4,4,3]
//	[4,4,2]
//	[4,4,1]
//	[4,3,4]
//	[4,3,3]
//	[4,3,2]
//	[4,3,1]
//	[4,2,4]
//	[4,2,3]
//	[4,2,2]
//	[4,2,1]
//	[4,1,4]
//	[4,1,3]
//	[4,1,2]
//	[4,1,1]
//	 ^ ^ ^
//	16 4 1 итерация 1 символа
//
//	alphabet: 1,2,3
//	length: 3
//	[3,3,3] 1 (0)
//	[3,3,2] 2 (1)
//	[3,3,1] 3 (2)
//	[3,2,3] 4 (3)
//	[3,2,2] 5 (4)
//	[3,2,1] 6 (5)
//	[3,1,3] 7 (6)
//	[3,1,2] 8 (7)
//	[3,1,1] 9 (8)
//	 ^ ^ ^
//	 9 3 1 итерация 1 символа
func DonePct(state []int, alphabetSize int) float64 {
	perChar := make([]uint64, len(state))
	for i := range perChar {
		a := len(perChar) - i - 1
		switch a {
		case 0:
			perChar[i] = 1
		case 1:
			perChar[i] = uint64(alphabetSize)
		default:
			perChar[i] = uint64(math.Pow(float64(alphabetSize), float64(a)))
		}
	}

	iterations := uint64(1)
	for i, cursor := range state {
		// +1 because of state has index position not the counter
		iterations += perChar[i] * uint64(alphabetSize-(cursor+1))
	}

	if iterations == 1 {
		return 0
	}

	return float64(iterations) / float64(Permutations(alphabetSize, len(state))) * 100
}
