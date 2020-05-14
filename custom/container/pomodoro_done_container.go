package container

import (
	"bitbucket.org/avanz/anotherPomodoro/common"
	"bitbucket.org/avanz/anotherPomodoro/custom/widget"
	"fmt"
	"fyne.io/fyne"
	"fyne.io/fyne/theme"
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
	var totalMinutes float64 = float64(time.Now().Hour()*60 + time.Now().Minute())
	timeDuration := 25 * time.Minute
	err := common.Settings.Read("settings", "timeDuration", &timeDuration)
	if err != nil {
		panic(err)
	}
	pauseDuration := 5 * time.Minute
	err = common.Settings.Read("settings", "pauseDuration", &pauseDuration)
	if err != nil {
		panic(err)
	}
	numberOfPresentPomodoro := int(math.Floor(totalMinutes / (timeDuration.Minutes() + pauseDuration.Minutes())))
	if len(c.Container.Objects) == numberOfPresentPomodoro+1 {
		return
	}

	if len(c.Container.Objects) == 0 {
		for i := 0; i < numberOfPresentPomodoro; i++ {
			pomodoro := widget.NewPomodoro(5, theme.BackgroundColor(), fmt.Sprintf("at %t", time.Now()))
			c.Container.AddObject(pomodoro)
		}
	}
	pomodoro := widget.NewPomodoro(5, common.Green, fmt.Sprintf("at %t", time.Now()))
	pomodoro.Show()
	c.Container.AddObject(pomodoro)
}
