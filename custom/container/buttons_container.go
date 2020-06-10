package container

import (
	"bitbucket.org/avanz/anotherPomodoro/custom/window"
	imgpause "bitbucket.org/avanz/anotherPomodoro/img/pause"
	imgstart "bitbucket.org/avanz/anotherPomodoro/img/start"
	"bitbucket.org/avanz/anotherPomodoro/repository"
	"bitbucket.org/avanz/anotherPomodoro/sync"
	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

type ButtonContainer struct {
	Container *fyne.Container
}

func NewButtonContainer(pause chan bool, syncRemoteAddressListener chan string, mainApp fyne.App, settingsRepository repository.IPomodoroRepository, synclistener sync.IListener) *fyne.Container {
	buttonsContainer := fyne.NewContainerWithLayout(
		layout.NewVBoxLayout(),
		NewPauseButton(pause, syncRemoteAddressListener),
		NewSettingsButton(syncRemoteAddressListener, mainApp, settingsRepository, synclistener))
	return buttonsContainer
}

func NewSettingsButton(syncRemoteAddressListener chan string, mainApp fyne.App, settingsRepository repository.IPomodoroRepository, synclistener sync.IListener) *widget.Button {
	settingsButton := widget.NewButton("Settings...", func() {
		settingsWindow := window.NewSettingsWindow(syncRemoteAddressListener, mainApp, settingsRepository, synclistener)
		settingsWindow.Show()
		settingsWindow.RequestFocus()
	})
	settingsButton.Resize(fyne.NewSize(30, 30))
	return settingsButton
}

func NewPauseButton(pause chan bool, syncRemoteAddressListener chan string) *widget.Button {
	var err error
	imgStart, err := imgstart.Asset("img/start/start.png")
	if err != nil {
		panic(err)
	}
	startIcon := fyne.NewStaticResource("start", imgStart)

	imgPause, err := imgpause.Asset("img/pause/pause.png")
	if err != nil {
		panic(err)
	}
	pauseIcon := fyne.NewStaticResource("pause", imgPause)

	icons := map[bool]fyne.Resource{false: startIcon, true: pauseIcon}

	var pauseFocusButton *widget.Button
	started := false
	pauseFocusButton = widget.NewButton("", func() {
		pause <- started
		pause <- started
		started = !started
		pauseFocusButton.SetIcon(icons[started])
	})

	go func(syncRemoteAddressListener chan string) {
		for {
			if <-syncRemoteAddressListener != "" {
				pauseFocusButton.Disable()
			} else {
				pauseFocusButton.Enable()
			}
		}
	}(syncRemoteAddressListener)
	pauseFocusButton.SetIcon(icons[started])
	pauseFocusButton.Resize(fyne.NewSize(30, 30))
	return pauseFocusButton
}
