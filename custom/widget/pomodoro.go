package widget

import (
	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"image/color"
)

type Pomodoro fyne.CanvasObject

func NewPomodoro(size int) Pomodoro {
	Pomodoro := canvas.NewRectangle(color.RGBA{R: 11, G: 156, B: 49, A: 1})
	Pomodoro.SetMinSize(fyne.Size{Width: size, Height: size})
	return Pomodoro
}
