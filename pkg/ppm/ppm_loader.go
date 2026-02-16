package ppm

import (
	"cmp"
	"image/color"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
)

type PPMImage struct {
	Width  int
	Height int
	Buffer []byte
}

func NewPPMImage(w, h int) PPMImage {
	return PPMImage{w, h, make([]byte, w*h*3)}
}

func OpenPPMP6(path string) (image PPMImage, err error) {
	file, err := os.Open(path)

	if err != nil {
		return NewPPMImage(0, 0), err
	}

	header_buf := make([]byte, 128)
	_, err = file.Read(header_buf)
	if err != nil {
		return NewPPMImage(0, 0), err
	}

	raw_header := string(header_buf)
	header_fields := strings.Fields(raw_header)

	w, err1 := strconv.Atoi(header_fields[1])
	h, err2 := strconv.Atoi(header_fields[2])
	if err := cmp.Or(err1, err2); err != nil {
		return NewPPMImage(0, 0), err
	}

	img := NewPPMImage(w, h)

	_, err = file.Seek(int64(-len(img.Buffer)), io.SeekEnd)
	if err != nil {
		return NewPPMImage(0, 0), err
	}

	_, err = file.Read(img.Buffer)
	if err != nil {
		return NewPPMImage(0, 0), err
	}

	return img, nil
}

func (img *PPMImage) CopyToEbitenImage() *ebiten.Image {
	ebit_img := ebiten.NewImage(img.Width, img.Height)
	for y := range img.Width {
		for x := range img.Height {
			idx := (y*img.Width + x) * 3
			clr := color.RGBA{img.Buffer[idx], img.Buffer[idx+1], img.Buffer[idx+2], 255}
			ebit_img.Set(x, y, clr)
		}
	}

	return ebit_img
}
