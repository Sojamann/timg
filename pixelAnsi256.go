package timg

import "image"

func rgbToAnsi256(r, g, b int) int {
	// convert the inclusive range 0-255 to 0-5
	conv := func(v int) int {
		switch {
		case v < 42:
			return 0
		case v < 84:
			return 1
		case v < 126:
			return 2
		case v < 168:
			return 3
		case v < 210:
			return 4
		default:
			return 5
		}
	}

	return (36 * conv(r)) + (6 * conv(g)) + (1 * conv(b)) + 16
}

func renderAnsi256(img image.Image) string {
	width := img.Bounds().Max.X - img.Bounds().Min.X
	height := img.Bounds().Max.Y - img.Bounds().Min.Y
	// TODO: different pre alloc amount
	result := make([]byte, 0, width*height)

	colStr := []byte("\x1B[38;5;000m")

	bounds := img.Bounds()
	for imgY := bounds.Min.Y; imgY < bounds.Max.Y; imgY += 2 {
		lastUpper, lastLower := -1, -1

		for imgX := bounds.Min.X; imgX < bounds.Max.X; imgX++ {
			curUpper := rgbToAnsi256(toRGB(img.At(imgX, imgY)))
			curLower := rgbToAnsi256(toRGB(img.At(imgX, imgY+1)))
			if curUpper == lastUpper && curLower == lastLower {
				result = append(result, "▀"...)
				continue
			}

			setAndPad(colStr[2:3], 3) // foreground
			setAndPad(colStr[7:10], curUpper)
			result = append(result, colStr...)

			if imgY+1 < bounds.Max.Y {
				setAndPad(colStr[2:3], 4)
				setAndPad(colStr[7:10], curLower)
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
