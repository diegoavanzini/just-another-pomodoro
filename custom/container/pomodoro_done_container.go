package container

import (
	"bitbucket.org/avanz/anotherPomodoro/custom/widget"
	"fmt"
	"fyne.io/fyne"
	"image/color"
	"math"
	"time"
)

type PomodoroDoneContainer struct {
	*fyne.Container
}

func NewPomodoroDoneContainer(boxLayout fyne.Layout) *PomodoroDoneContainer {
	container := &PomodoroDoneContainer{
		Container: fyne.NewContainerWithLayout(boxLayout),
	}
	return container
}

func (c PomodoroDoneContainer) AddPomodoro() {
	totalMinutes := time.Now().Hour()*60 + time.Now().Minute()
	numberOfPresentPomodoro := int(math.Floor(float64(totalMinutes / (25 + 5))))
	if len(c.Container.Objects) == numberOfPresentPomodoro+1 {
		return
	}
	green := color.RGBA{R: 11, G: 156, B: 49, A: 1}
	if len(c.Container.Objects) == 0 {
		for i := 0; i < numberOfPresentPomodoro; i++ {
			pomodoro := widget.NewPomodoro(5, green, fmt.Sprintf("at %t", time.Now()))
			pomodoro.Hide()
			c.Container.AddObject(pomodoro)
		}
	}
	pomodoro := widget.NewPomodoro(5, green, fmt.Sprintf("at %t", time.Now()))
	pomodoro.Show()
	c.Container.AddObject(pomodoro)
}
