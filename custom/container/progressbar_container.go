package container

import (
	"bitbucket.org/avanz/anotherPomodoro/common"
	custom_layout "bitbucket.org/avanz/anotherPomodoro/custom/layout"
	custom_widget "bitbucket.org/avanz/anotherPomodoro/custom/widget"
	"bitbucket.org/avanz/anotherPomodoro/repository"
	"bitbucket.org/avanz/anotherPomodoro/sync"
	"fyne.io/fyne"
	"time"
)

type ProgressBarContainer struct {
	Container *fyne.Container
}

func NewProgressBarContainer(pause, alert, addPomodoro chan bool, syncRemoteAddressListener chan string, pomodoroRepository repository.IPomodoroRepository) ProgressBarContainer {
	var timerDuration, pauseDuration time.Duration
	err := pomodoroRepository.Read("settings", "timeDuration", &timerDuration)
	if err != nil {
		err = pomodoroRepository.Write("settings", "timeDuration", 25*time.Minute)
		if err != nil {
			panic(err)
		}
	}
	err = pomodoroRepository.Read("settings", "pauseDuration", &pauseDuration)
	if err != nil {
		err = pomodoroRepository.Write("settings", "pauseDuration", 5*time.Minute)
		if err != nil {
			panic(err)
		}
	}

	pomodoroRepository.Write("settings", "synkAddress", "") // reset synk remote

	progressBar := custom_widget.NewTimerProgressBar(timerDuration, pause, alert, common.Yellow, pomodoroRepository, "timer")
	progressBar.Resize(fyne.NewSize(220, 5))
	pauseProgressBar := custom_widget.NewTimerProgressBar(pauseDuration, pause, alert, common.Green, pomodoroRepository, "pause")
	pauseProgressBar.Resize(fyne.NewSize(50, 5))
	progressBarContainer := fyne.NewContainerWithLayout(
		custom_layout.NewResizableGridLayout(2),
		progressBar,
		pauseProgressBar)


	go func(syncRemoteAddressListener chan string) {
		for {
			remoteAddress := <-syncRemoteAddressListener
			if remoteAddress == "" {
				progressBar.SetSyncClient(nil)
			} else {
				progressBar.SetSyncClient(sync.NewTcpClient(remoteAddress))
			}
		}
	}(syncRemoteAddressListener)

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
