package container

import (
	custom_layout "bitbucket.org/avanz/anotherPomodoro/custom/layout"
	custom_widget "bitbucket.org/avanz/anotherPomodoro/custom/widget"
	"fyne.io/fyne"
	"time"
)

type ProgressBarContainer struct {
	Container *fyne.Container
}

func NewProgressBarContainer(pause, alert, addPomodoro chan bool) ProgressBarContainer {
	progressBar := custom_widget.NewTimerProgressBar(25*time.Minute, pause, alert)
	progressBar.Resize(fyne.NewSize(250, 5))
	pauseProgressBar := custom_widget.NewTimerProgressBar(5*time.Minute, pause, alert)
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
