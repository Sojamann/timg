package timg

import (
	"image"
	"os"

	"github.com/muesli/termenv"
	"golang.org/x/term"
)

type opt struct {
	width, height int
	profile       termenv.Profile
}

type RenderOption func(o *opt)

// FitTo is an option to make the image fit into the box
// defined by width and height
func FitTo(width, height int) RenderOption {
	return func(o *opt) {
		o.width = width
		o.height = height
	}
}

// FitToScreen tries to figure out the dimentions of the
// terminal (on stdout). If this fails ... no image dimensions
// are set. This is so that the image rendering can never fail.
// If you are unsure about this behavior just use FitTo.
func FitScreen() RenderOption {
	return func(o *opt) {
		width, height, err := term.GetSize(1)
		if err != nil {
			return
		}
		o.width = width
		o.height = height
	}
}

// WithColorProfile is an option that makes Render use
// the given profile instead of querying the terminal(emulator)
// for what it supports.
func WithColorProfile(profile termenv.Profile) RenderOption {
	return func(o *opt) {
		o.profile = profile
	}
}

// Render returns the string representing the rendered image taking
// the RenderOptions into account.
func Render(img image.Image, opts ...RenderOption) string {
	renderer := opt{profile: -1}
	for _, opt := range opts {
		opt(&renderer)
	}

	if renderer.width == 0 && renderer.height == 0 {
		bounds := img.Bounds()
		renderer.width = bounds.Max.X - bounds.Min.X
		renderer.height = bounds.Max.Y - bounds.Min.Y
	}

	if renderer.profile == -1 {
		renderer.profile = termenv.NewOutput(os.Stdout).Profile
	}

	img = resizeImg(img, renderer.width, renderer.height)

	switch renderer.profile {
	case termenv.ANSI:
		return renderAnsi(img)
	case termenv.ANSI256:
		return renderAnsi256(img)
	case termenv.TrueColor:
		return renderTrueColor(img)
	default: // default & ascii profile
		return renderAscii(img)
	}
}
