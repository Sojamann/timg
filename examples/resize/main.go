package main

import (
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"

	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/termenv"
	"github.com/sojamann/timg"
)

func die(msg string) {
	os.Stderr.WriteString(msg)
	os.Exit(1)
}

func main() {
	var img image.Image

	// load the image passed via cli arg
	if len(os.Args) < 2 {
		die(fmt.Sprintf("Usage: %s <path/to/imgage>", os.Args[0]))
	}
	fp, err := os.Open(os.Args[1])
	if err != nil {
		die(err.Error())
	}
	defer fp.Close()

	img, _, err = image.Decode(fp)
	if err != nil {
		die(err.Error())
	}

	ascii := timg.Render(img, timg.FitTo(15, 15), timg.WithColorProfile(termenv.Ascii))
	ansi := timg.Render(img, timg.FitTo(15, 15), timg.WithColorProfile(termenv.ANSI))
	ansi256 := timg.Render(img, timg.FitTo(15, 15), timg.WithColorProfile(termenv.ANSI256))
	trueColor := timg.Render(img, timg.FitTo(15, 15), timg.WithColorProfile(termenv.TrueColor))

	fmt.Println(lipgloss.JoinHorizontal(lipgloss.Center, ascii, ansi, ansi256, trueColor))
}
