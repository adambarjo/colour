package colour

import (
	"errors"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

var (
	errInvalidColourFmt = errors.New("invalid colour format")
)

type ColourFmt int

const (
	Hex ColourFmt = iota
	Rgb
	Rgba
	Srgb
)

type Colour struct {
	F       ColourFmt
	R, G, B int
	A       float64
}

func (c *Colour) String() string {
	return fmt.Sprintf(
		"F:%d (%s), Hex: %s, Rgb: %s, Rgba: %s, Srgb: %s",
		c.F, map[ColourFmt]string{
			Hex:  "hex",
			Rgb:  "rgb",
			Rgba: "rgba",
			Srgb: "srgb",
		}[c.F], c.To(Hex), c.To(Rgb), c.To(Rgba), c.To(Srgb),
	)
}

// Parse is used to parse a string to a `Colour`
// the string can be any of the following formats:
// rgb(1, 1, 1)
// rgba(1, 1, 1, 0.5)
// #ff00cc
// #ff00cc80
// color(srgb 1, 1, 1)
func Parse(s string) (c *Colour, err error) {
	s = strings.ToLower(strings.TrimSpace(s))
	re := regexp.MustCompile(
		`rgb\s*\(\s*\d+\s*,\s*\d+\s*,\s*\d+\s*\)` +
			`|rgba\s*\(\s*\d+\s*,\s*\d+\s*,\s*\d+\s*,\s*(0|1|0\.\d+)\s*\)` +
			`|#[a-fA-F\d]{6,8}` +
			`|color\s*\(\s*srgb(\s+(0|1)(\.\d+)?){3}\s*\)`,
	)

	if !re.Match([]byte(s)) {
		return nil, errInvalidColourFmt
	}

	if s[0] == '#' {
		values, err := strconv.ParseUint(s[1:7], 16, 32)
		if err != nil {
			return nil, err
		}

		c = &Colour{
			F: Hex,
			R: int(values >> 16),
			G: int((values >> 8) & 0xFF),
			B: int(values & 0xFF),
			A: 1,
		}

		if len(s) == 9 {
			alpha, err := strconv.ParseUint(s[7:], 16, 32)
			if err != nil {
				return nil, err
			}
			c.A = float64(alpha&0xFF) / 255.0
		}

		return c, nil
	}

	if strings.HasPrefix(s, "rgb") {
		s = strings.TrimPrefix(s, "rgba")
		s = strings.TrimPrefix(s, "rgb") // (255, 255, 255)
		s = strings.TrimSpace(s)
		s = s[1 : len(s)-1] // 255, 255, 255

		values, err := stringsToFloats(strings.Split(s, ","), 64)
		if err != nil {
			return nil, errInvalidColourFmt
		}

		c = &Colour{
			F: Rgb,
			R: int(values[0]),
			G: int(values[1]),
			B: int(values[2]),
			A: 1,
		}

		if len(values) == 4 {
			c.F = Rgba
			c.A = values[3]
		}

		return c, nil
	}

	if strings.HasPrefix(s, "color") {
		s = strings.TrimPrefix(s, "color") // (srgb 1 1 1)
		s = strings.TrimSpace(s)
		s = s[1 : len(s)-1] // srgb 1 1 1
		s = strings.TrimSpace(s)
		s = strings.TrimPrefix(s, "srgb") // 1 1 1
		s = strings.TrimSpace(s)

		values, err := stringsToFloats(strings.Split(s, " "), 64)
		if err != nil {
			return nil, err
		}

		c = &Colour{
			F: Srgb,
			R: int(math.Round(255 * values[0])),
			G: int(math.Round(255 * values[1])),
			B: int(math.Round(255 * values[2])),
			A: 1,
		}

		return c, nil
	}

	return nil, errInvalidColourFmt
}

// To is used to convert a *Colour to a specified colour format
func (c *Colour) To(colourFmt ColourFmt) string {
	switch colourFmt {
	case Hex:
		if c.A == 1 {
			return fmt.Sprintf("#%02x%02x%02x", c.R, c.G, c.B)
		}
		return fmt.Sprintf("#%02x%02x%02x%02x", c.R, c.G, c.B, int(math.Round(255*c.A)))
	case Rgb:
		return fmt.Sprintf("rgb(%d, %d, %d)", c.R, c.G, c.B)
	case Rgba:
		return fmt.Sprintf("rgba(%d, %d, %d, %.2f)", c.R, c.G, c.B, c.A)
	case Srgb:
		return fmt.Sprintf("color(srgb %.6f %.6f %.6f)", float64(c.R)/255, float64(c.G)/255, float64(c.B)/255)
	default:
		return ""
	}
}
