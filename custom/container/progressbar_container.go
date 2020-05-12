package container

import (
	"bitbucket.org/avanz/anotherPomodoro/common"
	custom_layout "bitbucket.org/avanz/anotherPomodoro/custom/layout"
	custom_widget "bitbucket.org/avanz/anotherPomodoro/custom/widget"
	"fyne.io/fyne"
	"time"
)

type ProgressBarContainer struct {
	Container *fyne.Container
}

func NewProgressBarContainer(pause, alert, addPomodoro chan bool) ProgressBarContainer {
	var timerDuration, pauseDuration time.Duration
	err := common.Settings.Read("settings", "timeDuration", &timerDuration)
	if err != nil {
		err = common.Settings.Write("settings", "timeDuration", 25*time.Minute)
		if err != nil {
			panic(err)
		}
	}
	err = common.Settings.Read("settings", "pauseDuration", &timerDuration)
	if err != nil {
		err = common.Settings.Write("settings", "pauseDuration", 5*time.Minute)
		if err != nil {
			panic(err)
		}
	}
	progressBar := custom_widget.NewTimerProgressBar(timerDuration, pause, alert)
	progressBar.Resize(fyne.NewSize(250, 5))
	pauseProgressBar := custom_widget.NewTimerProgressBar(pauseDuration, pause, alert)
	pauseProgressBar.Resize(fyne.NewSize(50, 5))
	progressBarContainer := fyne.NewContainerWithLayout(
		custom_layout.NewResizableGridLayout(2),
		progressBar,
		pauseProgressBar)
	go func() {
		for {
			pauseProgressBar.SetValue(time.Duration(0))
			progressBar.Start()
			pauseProgressBar.Start()
			addPomodoro <- true
		}
	}()
	return ProgressBarContainer{progressBarContainer}
}
