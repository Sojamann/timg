package timg

import (
	"fmt"
	"image"
	"testing"
)

func TestResizeImg(t *testing.T) {
	type size struct {
		width, height int
	}

	type testcase struct {
		from     size
		to       size
		expected size
	}

	testcases := []testcase{
		{size{200, 100}, size{50, 25}, size{50, 25}},
		{size{200, 100}, size{25, 50}, size{25, 12}},
		{size{100, 200}, size{50, 25}, size{12, 25}},
		{size{100, 100}, size{50, 50}, size{50, 50}},
	}

	for _, testcase := range testcases {
		t.Run(fmt.Sprintf("%v -> %v", testcase.from, testcase.to), func(t *testing.T) {
			input := image.NewRGBA(image.Rect(0, 0, testcase.from.width, testcase.from.height))

			res := resizeImg(input, testcase.to.width, testcase.to.height)

			bounds := res.Bounds()
			actualWidth := bounds.Max.X - bounds.Min.X
			actualHeight := bounds.Max.Y - bounds.Min.Y

			if actualWidth != testcase.expected.width {
				t.Errorf("image has width %d but %d was expected", actualWidth, testcase.expected.width)
			}
			if actualHeight != testcase.expected.height {
				t.Errorf("image has height %d but %d was expected", actualHeight, testcase.expected.height)
			}
		})
	}
}
