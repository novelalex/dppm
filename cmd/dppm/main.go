package main

import (
	"flag"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/novelalex/dppm/pkg/ppm"
)

func main() {
	file_path := flag.String("file", "img.ppm", "Path to image")
	flag.Parse()

	img, err := ppm.OpenPPMP6(*file_path)
	if err != nil {
		log.Fatal(err)
	}

	ebit_img := img.CopyToEbitenImage()
	state := State{
		X: 0, Y: 0, Zoom: 0, MouseDown: false,
	}
	app := App{
		*file_path,
		img.Width,
		img.Height,
		ebit_img,
		&state,
	}

	ebiten.SetWindowSize(1280, 720)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowTitle(*file_path)
	if err := ebiten.RunGame(&app); err != nil {
		log.Fatal(err)
	}

}
