package timg

import "image"

func renderTrueColor(img image.Image) string {
	width := img.Bounds().Max.X - img.Bounds().Min.X
	height := img.Bounds().Max.Y - img.Bounds().Min.Y
	// TODO: different pre alloc amount
	result := make([]byte, 0, width*height)

	colStr := []byte("\x1B[38;2;000;000;000m")

	bounds := img.Bounds()
	for imgY := bounds.Min.Y; imgY < bounds.Max.Y; imgY += 2 {
		lastUpperR, lastUpperG, lastUpperB := -1, -1, -1
		lastLowerR, lastLowerG, lastLowerB := -1, -1, -1

		for imgX := bounds.Min.X; imgX < bounds.Max.X; imgX++ {
			upperR, upperG, upperB := toRGB(img.At(imgX, imgY))
			lowerR, lowerG, lowerB := toRGB(img.At(imgX, imgY+1))
			if lastUpperR == upperR &&
				lastUpperG == upperG &&
				lastUpperB == upperB &&
				lastLowerR == lowerR &&
				lastLowerG == lowerG &&
				lastLowerB == lowerB {
				result = append(result, "▀"...)
				continue
			}

			setAndPad(colStr[2:3], 3) // foreground
			setAndPad(colStr[7:10], upperR)
			setAndPad(colStr[11:14], upperG)
			setAndPad(colStr[15:18], upperB)
			result = append(result, colStr...)

			if imgY+1 < bounds.Max.Y {
				setAndPad(colStr[2:3], 4) // background
				setAndPad(colStr[7:10], lowerR)
				setAndPad(colStr[11:14], lowerG)
				setAndPad(colStr[15:18], lowerB)
				result = append(result, colStr...)
			} else {
				// for uneven number of rows we want the last row
				// to use the default color of the terminal so that it
				// is invisible
				result = append(result, "\x1B[49m"...)
			}

			result = append(result, "▀"...)

			lastUpperR, lastUpperG, lastUpperB = upperR, upperG, upperB
			lastLowerR, lastLowerG, lastLowerB = lowerR, lowerG, lowerB
		}
		// there must be a reset on every line
		result = append(result, "\x1B[0m\n"...)
	}

	return string(result)
}
