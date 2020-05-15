package widget

import (
	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"image/color"
)

type Pomodoro fyne.CanvasObject

func NewPomodoro(size int, color color.Color) Pomodoro {
	pomodoro := canvas.NewRectangle(color)
	pomodoro.SetMinSize(fyne.Size{Width: 2, Height: size})
	return pomodoro
}
