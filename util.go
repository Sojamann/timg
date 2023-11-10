package timg

import (
	"image"
	"image/color"
	"strconv"

	"golang.org/x/image/draw"
)

func setAndPad(target []byte, value int) {
	conv := strconv.Itoa(value)
	offset := len(target) - len(conv)
	for i := 0; i < offset; i++ {
		target[i] = '0'
	}
	copy(target[offset:], conv)
}

func toRGB(c color.Color) (int, int, int) {
	r, g, b, _ := c.RGBA()
	return int(uint8(r >> 8)), int(uint8(g >> 8)), int(uint8(b >> 8))
}

// resizeImg resizes an image to fit in the box defiend by
// boxWidth and boxHeight whilst keeping it's current aspect ratio
func resizeImg(img image.Image, boxWidth, boxHeight int) image.Image {
	bounds := img.Bounds()
	actualWidth := bounds.Max.X - bounds.Min.X
	actualHeight := bounds.Max.Y - bounds.Min.Y

	ratio := float64(actualWidth) / float64(actualHeight)

	var resultingWidth, resultingHeight int
	if actualWidth > actualHeight {
		resultingWidth = boxWidth
		resultingHeight = int(float64(boxWidth) / ratio)
	} else {
		resultingHeight = boxHeight
		resultingWidth = int(float64(boxHeight) * ratio)
	}

	result := image.NewRGBA(image.Rect(0, 0, resultingWidth, resultingHeight))
	draw.NearestNeighbor.Scale(result, result.Bounds(), img, img.Bounds(), draw.Over, nil)

	return result
}
