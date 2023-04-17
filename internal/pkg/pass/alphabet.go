package pass

import (
	"math/rand"
	"strings"

	"github.com/samber/lo"
)

var Digits = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
var Symbols = []string{"`", `"`, "~", "!", "@", "#", "$", "%", "^", "&", "*", "(", ")", "–", "-", " ", "_", "=", "+", "[", "]", "{", "}", "\\", "|", ";", ":", "‘", "“", ",", ".", "/", "<", ">", "?"}
var EnglishLower = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
var EnglishUpper = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

var Dictionary = map[string][]string{
	"digits":        Digits,
	"symbols":       Symbols,
	"english-lower": EnglishLower,
	"english-upper": EnglishUpper,
}

func ParseAlphabets(in []string) []string {
	out := make([]string, 0)
	for _, v := range in {
		dic, ok := Dictionary[v]
		switch {
		case ok:
			out = append(out, dic...)
		default:
			out = append(out, strings.Split(v, "")...)
		}
	}

	return lo.Uniq(out)
}

func ShuffleIfNonZero(in []string, seed int64) []string {
	if seed == 0 {
		return in
	}

	out := append(make([]string, 0), in...)
	r := rand.New(rand.NewSource(seed))
	r.Shuffle(len(out), func(i, j int) {
		out[i], out[j] = out[j], out[i]
	})

	return out
}
