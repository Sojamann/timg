package timg

import (
	"image"

	"github.com/lucasb-eyer/go-colorful"
)

// These colors will only be used for comparison in order
//
//	to choose the right ansi color code from given rgb values.
//
// Note: do not re-order as the index corresponds to the ansi code.
var ansiColors = []colorful.Color{
	colorful.LinearRgb(0.0/256, 0.0/256, 0.0/256),       // black
	colorful.LinearRgb(205.0/256, 49.0/256, 49.0/256),   // red
	colorful.LinearRgb(13.0/256, 188.0/256, 121.0/256),  // green
	colorful.LinearRgb(229.0/256, 229.0/256, 16.0/256),  // yellow
	colorful.LinearRgb(36.0/256, 114.0/256, 200.0/256),  // blue
	colorful.LinearRgb(188.0/256, 63.0/256, 188.0/256),  // magenta
	colorful.LinearRgb(17.0/256, 168.0/256, 205.0/256),  // cyan
	colorful.LinearRgb(229.0/256, 229.0/256, 229.0/256), // white
}

func rgbToLab(r, g, b int) colorful.Color {
	x, y, z := colorful.LinearRgbToXyz(float64(r)/256, float64(g)/256, float64(b)/256)
	l, a, b_ := colorful.XyzToLab(x, y, z)
	return colorful.Lab(l, a, b_)
}

func rgbToAnsi(r, g, b int) int {
	asLab := rgbToLab(r, g, b)

	closestIndex, closestDist := 0, 999999.0
	for i, ansiCol := range ansiColors {
		dist := asLab.DistanceCIE76(ansiCol)

		if dist < closestDist {
			closestIndex = i
			closestDist = dist
		}
	}

	return closestIndex
}

func renderAnsi(img image.Image) string {
	width := img.Bounds().Max.X - img.Bounds().Min.X
	height := img.Bounds().Max.Y - img.Bounds().Min.Y
	// TODO: different pre alloc amount
	result := make([]byte, 0, width*height)

	colStr := []byte("\x1B[30m")

	bounds := img.Bounds()
	for imgY := bounds.Min.Y; imgY < bounds.Max.Y; imgY += 2 {
		lastUpper, lastLower := -1, -1

		for imgX := bounds.Min.X; imgX < bounds.Max.X; imgX++ {
			curUpper := rgbToAnsi(toRGB(img.At(imgX, imgY)))
			curLower := rgbToAnsi(toRGB(img.At(imgX, imgY+1)))
			if curUpper == lastUpper && curLower == lastLower {
				result = append(result, "▀"...)
				continue
			}

			setAndPad(colStr[2:3], 3) // foreground
			setAndPad(colStr[3:4], curUpper)
			result = append(result, colStr...)

			if imgY+1 < bounds.Max.Y {
				setAndPad(colStr[2:3], 4)
				setAndPad(colStr[3:4], curLower)
				result = append(result, colStr...)
			} else {
				// for uneven number of rows we want the last row
				// to use the default color of the terminal so that it
				// is invisible
				result = append(result, "\x1B[49m"...)
			}

			result = append(result, "▀"...)
			lastUpper, lastLower = curUpper, curLower
		}
		// there must be a reset on every line
		result = append(result, "\x1B[0m\n"...)
	}

	return string(result)
}
