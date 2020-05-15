package container

import (
	"bitbucket.org/avanz/anotherPomodoro/common"
	custom_layout "bitbucket.org/avanz/anotherPomodoro/custom/layout"
	custom_widget "bitbucket.org/avanz/anotherPomodoro/custom/widget"
	"bitbucket.org/avanz/anotherPomodoro/repository"
	"fyne.io/fyne"
	"time"
)

type ProgressBarContainer struct {
	Container *fyne.Container
}

func NewProgressBarContainer(pause, alert, addPomodoro chan bool, repositorySettings repository.ISettingsRepository) ProgressBarContainer {
	var timerDuration, pauseDuration time.Duration
	err := repositorySettings.Read("settings", "timeDuration", &timerDuration)
	if err != nil {
		err = repositorySettings.Write("settings", "timeDuration", 25*time.Minute)
		if err != nil {
			panic(err)
		}
	}
	err = repositorySettings.Read("settings", "pauseDuration", &pauseDuration)
	if err != nil {
		err = repositorySettings.Write("settings", "pauseDuration", 5*time.Minute)
		if err != nil {
			panic(err)
		}
	}
	progressBar := custom_widget.NewTimerProgressBar(timerDuration, pause, alert, common.Red)
	progressBar.Resize(fyne.NewSize(220, 5))
	pauseProgressBar := custom_widget.NewTimerProgressBar(pauseDuration, pause, alert, common.Green)
	pauseProgressBar.Resize(fyne.NewSize(50, 5))
	progressBarContainer := fyne.NewContainerWithLayout(
		custom_layout.NewResizableGridLayout(2),
		progressBar,
		pauseProgressBar)
	go func() {
		for {
			pauseProgressBar.SetValue(pauseDuration)
			progressBar.Start()
			pauseProgressBar.Start()
			addPomodoro <- true
		}
	}()
	return ProgressBarContainer{progressBarContainer}
}
