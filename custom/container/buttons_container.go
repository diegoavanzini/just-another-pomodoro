package container

import (
	"bitbucket.org/avanz/anotherPomodoro/custom/window"
	"bitbucket.org/avanz/anotherPomodoro/repository"
	"bitbucket.org/avanz/anotherPomodoro/sync"
	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	packr "github.com/gobuffalo/packr/v2"
	"os"
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
	startIcon, err := fyne.LoadResourceFromPath(packr.New("start", "."+string(os.PathSeparator)+"img"+string(os.PathSeparator)+"start.png").ResolutionDir)
	if err != nil {
		panic(err)
	}
	pauseIcon, err := fyne.LoadResourceFromPath(packr.New("pause", "."+string(os.PathSeparator)+"img"+string(os.PathSeparator)+"pause.png").ResolutionDir)
	if err != nil {
		panic(err)
	}
	icons := map[bool]fyne.Resource{false: startIcon, true: pauseIcon}

	var pauseFocusButton *widget.Button
	toPause := true
	pauseFocusButton = widget.NewButton("", func() {
		pause <- toPause
		pause <- toPause
		toPause = !toPause
		pauseFocusButton.SetIcon(icons[toPause])
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
	pauseFocusButton.SetIcon(icons[toPause])
	pauseFocusButton.Resize(fyne.NewSize(30, 30))
	return pauseFocusButton
}
