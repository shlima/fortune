package pass

import (
	"fmt"
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
	return strings.Join(in, ",")
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
