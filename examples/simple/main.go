package main

import (
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"

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

	// render and print the image in whatever size it is in right now
	fmt.Println(timg.Render(img))
}
