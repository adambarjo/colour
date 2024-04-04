# Colour

Convert colour formats from one to another.

## Installation

`go get github.com/vqvw/colour`

## Usage

```go
myColour := colour.Parse("#ff0000")
myColour.To(colour.Hex)  // #ff0000
myColour.To(colour.Rgb)  // rgb(255, 0, 0)
myColour.To(colour.Rgba) // rgb(255, 0, 0, 1)
myColour.To(colour.Srgb) // color(srgb 1.000000 0.000000 0.000000)
```
