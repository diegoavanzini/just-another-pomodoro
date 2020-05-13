package widget

import (
	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"image/color"
)

type Pomodoro fyne.CanvasObject

func NewPomodoro(size int, color color.Color, tooltip string) Pomodoro {
	pomodoro := canvas.NewRectangle(color)
	pomodoro.SetMinSize(fyne.Size{Width: 1, Height: size})
	return pomodoro
}
