package container

import (
	"bitbucket.org/avanz/anotherPomodoro/custom/widget"
	"fyne.io/fyne"
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
	if len(c.Container.Objects) == 0 {
		for i := 0; i < numberOfPresentPomodoro; i++ {
			pomodoro := domain.NewPomodoro(5)
			pomodoro.Hide()
			c.Container.AddObject(pomodoro)
		}
	}
	pomodoro := domain.NewPomodoro(5)
	pomodoro.Show()
	c.Container.AddObject(pomodoro)
}
