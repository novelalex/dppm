package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type State struct {
	X, Y             float64
	OffsetX, OffsetY float64
	Zoom             float64
	MouseDown        bool
	MouseX           float64
	MouseY           float64
}

type App struct {
	FileName      string
	Width, Height int
	Image         *ebiten.Image
	State         *State
}

func (a *App) Update() error {

	mX, mY := ebiten.CursorPosition()
	a.State.MouseX, a.State.MouseY = float64(mX), float64(mY)

	_, sY := ebiten.Wheel()

	a.State.Zoom += sY

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		a.State.MouseDown = true
		a.State.OffsetX = a.State.X - float64(mX)
		a.State.OffsetY = a.State.Y - float64(mY)

	}

	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		a.State.MouseDown = false
	}

	if a.State.MouseDown {
		a.State.X = float64(mX) + a.State.OffsetX
		a.State.Y = float64(mY) + a.State.OffsetY
	}

	return nil
}

func (a *App) Draw(screen *ebiten.Image) {
	op := ebiten.DrawImageOptions{}
	w := float64(a.Image.Bounds().Bounds().Dx())
	h := float64(a.Image.Bounds().Bounds().Dy())
	window_w, window_h := ebiten.WindowSize()
	win_center_x := float64(window_w) / 2
	win_center_y := float64(window_h) / 2

	op.GeoM.Translate(-w/2, -h/2)
	op.GeoM.Scale(1+a.State.Zoom/10, 1+a.State.Zoom/10)
	op.GeoM.Translate(win_center_x, win_center_y)
	op.GeoM.Translate(a.State.X, a.State.Y)
	screen.DrawImage(a.Image, &op)

	display_str := fmt.Sprintf("%s\n%dx%d", a.FileName, a.Width, a.Height)
	ebitenutil.DebugPrint(screen, display_str)

}

func (a *App) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}
