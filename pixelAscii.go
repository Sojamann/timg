package timg

import (
	"image"
	"image/color"
)

// every pixel is two wide
func rgbToAscii(r, g, b int) string {
	col := color.RGBA{uint8(r), uint8(g), uint8(b), 0}
	gray := color.GrayModel.Convert(col).(color.Gray)

	if gray.Y > 170 {
		return "██"
	}
	return "  "
}

func renderAscii(img image.Image) string {
	bounds := img.Bounds()
	width := bounds.Max.X - bounds.Min.X
	height := bounds.Max.Y - bounds.Min.Y

	// img size plus a new line on every row
	result := make([]byte, 0, ((width*2)*height)+height)

	for imgY := bounds.Min.Y; imgY < bounds.Max.Y; imgY++ {
		for imgX := bounds.Min.X; imgX < bounds.Max.X; imgX++ {
			char := rgbToAscii(toRGB(img.At(imgX, imgY)))
			result = append(result, char...)
		}
		result = append(result, '\n')
	}

	return string(result)
}
