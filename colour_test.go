package colour

import (
	"testing"
)

type parseTest struct {
	input, expected string
}

var parseTests = []parseTest{
	{
		"#ff00cc",
		"F:0 (hex), Hex: #ff00cc, Rgb: rgb(255, 0, 204), Rgba: rgba(255, 0, 204, 1.00), Srgb: color(srgb 1.000000 0.000000 0.800000)",
	},
	{
		"#ff00cc80",
		"F:0 (hex), Hex: #ff00cc80, Rgb: rgb(255, 0, 204), Rgba: rgba(255, 0, 204, 0.50), Srgb: color(srgb 1.000000 0.000000 0.800000)",
	},
	{
		"#ff00cc40",
		"F:0 (hex), Hex: #ff00cc40, Rgb: rgb(255, 0, 204), Rgba: rgba(255, 0, 204, 0.25), Srgb: color(srgb 1.000000 0.000000 0.800000)",
	},
	{
		"rgb(255, 0, 204)",
		"F:1 (rgb), Hex: #ff00cc, Rgb: rgb(255, 0, 204), Rgba: rgba(255, 0, 204, 1.00), Srgb: color(srgb 1.000000 0.000000 0.800000)",
	},
	{
		"rgb(255,0,204)",
		"F:1 (rgb), Hex: #ff00cc, Rgb: rgb(255, 0, 204), Rgba: rgba(255, 0, 204, 1.00), Srgb: color(srgb 1.000000 0.000000 0.800000)",
	},
	{
		"rgb   (   255   ,   0   ,   204   )",
		"F:1 (rgb), Hex: #ff00cc, Rgb: rgb(255, 0, 204), Rgba: rgba(255, 0, 204, 1.00), Srgb: color(srgb 1.000000 0.000000 0.800000)",
	},
	{
		"rgba(255, 0, 204, 1)",
		"F:2 (rgba), Hex: #ff00cc, Rgb: rgb(255, 0, 204), Rgba: rgba(255, 0, 204, 1.00), Srgb: color(srgb 1.000000 0.000000 0.800000)",
	},
	{
		"rgba(255, 0, 204, 0.5)",
		"F:2 (rgba), Hex: #ff00cc80, Rgb: rgb(255, 0, 204), Rgba: rgba(255, 0, 204, 0.50), Srgb: color(srgb 1.000000 0.000000 0.800000)",
	},
	{
		"rgba(255,0,204,0.5)",
		"F:2 (rgba), Hex: #ff00cc80, Rgb: rgb(255, 0, 204), Rgba: rgba(255, 0, 204, 0.50), Srgb: color(srgb 1.000000 0.000000 0.800000)",
	},
	{
		"rgba   (   255   ,   0   ,   204   ,   0.5   )",
		"F:2 (rgba), Hex: #ff00cc80, Rgb: rgb(255, 0, 204), Rgba: rgba(255, 0, 204, 0.50), Srgb: color(srgb 1.000000 0.000000 0.800000)",
	},
	{
		"color(srgb 1.00000 0.00000 0.80000)",
		"F:3 (srgb), Hex: #ff00cc, Rgb: rgb(255, 0, 204), Rgba: rgba(255, 0, 204, 1.00), Srgb: color(srgb 1.000000 0.000000 0.800000)",
	},
	{
		"#ebf2ff",
		"F:0 (hex), Hex: #ebf2ff, Rgb: rgb(235, 242, 255), Rgba: rgba(235, 242, 255, 1.00), Srgb: color(srgb 0.921569 0.949020 1.000000)",
	},
	{
		"color(srgb 0.921569 0.949020 1)",
		"F:3 (srgb), Hex: #ebf2ff, Rgb: rgb(235, 242, 255), Rgba: rgba(235, 242, 255, 1.00), Srgb: color(srgb 0.921569 0.949020 1.000000)",
	},
	{
		"color   (   srgb   0.921569   0.949020   1   )",
		"F:3 (srgb), Hex: #ebf2ff, Rgb: rgb(235, 242, 255), Rgba: rgba(235, 242, 255, 1.00), Srgb: color(srgb 0.921569 0.949020 1.000000)",
	},
}

func TestParse(t *testing.T) {
	for _, test := range parseTests {
		output, err := Parse(test.input)
		if outputString := output.String(); outputString != test.expected {
			t.Errorf(`
        expected: %s
        received: %s`, test.expected, outputString)
		}
		if err != nil {
			t.Error("error:", err)
		}
	}
}

type toTest struct {
	colourInput string
	colourFmt   ColourFmt
	expected    string
}

var toTests = []toTest{
	{
		"#ff00cc",
		Hex,
		"#ff00cc",
	},
	{
		"#ff00cc80",
		Hex,
		"#ff00cc80",
	},
	{
		"#ff00cc",
		Rgb,
		"rgb(255, 0, 204)",
	},
	{
		"#ff00cc80",
		Rgb,
		"rgb(255, 0, 204)",
	},
	{
		"#ff00cc",
		Rgba,
		"rgba(255, 0, 204, 1.00)",
	},
	{
		"#ff00cc80",
		Rgba,
		"rgba(255, 0, 204, 0.50)",
	},
	{
		"#ff00cc",
		Srgb,
		"color(srgb 1.000000 0.000000 0.800000)",
	},
	{
		"#ff00cc80",
		Srgb,
		"color(srgb 1.000000 0.000000 0.800000)",
	},
}

func TestTo(t *testing.T) {
	for _, test := range toTests {
		myColour, _ := Parse(test.colourInput)
		output := myColour.To(test.colourFmt)
		if output != test.expected {
			t.Errorf(`
        expected: %s
        received: %s`, test.expected, output)
		}
	}
}
