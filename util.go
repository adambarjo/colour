package colour

import (
	"strconv"
	"strings"
)

// stringsToFloats converts a slice of strings to a slice of floats
func stringsToFloats(strSlice []string, bitSize int) (out []float64, err error) {
	for _, s := range strSlice {
		s = strings.TrimSpace(s)
		if s == "" {
			continue
		}
		f, err := strconv.ParseFloat(s, bitSize)
		if err != nil {
			return nil, err
		}
		out = append(out, f)
	}
	return
}
