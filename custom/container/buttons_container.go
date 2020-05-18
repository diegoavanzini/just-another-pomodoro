package container

import (
	"bitbucket.org/avanz/anotherPomodoro/custom/window"
	"bitbucket.org/avanz/anotherPomodoro/repository"
	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"github.com/gobuffalo/packr/v2"
	"os"
)

type ButtonContainer struct {
	Container *fyne.Container
}

func NewButtonContainer(pause chan bool, inSync chan string, mainApp fyne.App, settingsRepository repository.IPomodoroRepository) *fyne.Container {
	buttonsContainer := fyne.NewContainerWithLayout(
		layout.NewVBoxLayout(),
		NewPauseButton(pause, inSync),
		NewSettingsButton(inSync, mainApp, settingsRepository))
	return buttonsContainer
}

func NewSettingsButton(inSync chan string, mainApp fyne.App, settingsRepository repository.IPomodoroRepository) *widget.Button {
	settingsButton := widget.NewButton("Settings...", func() {
		settingsWindow := window.NewSettingsWindow(inSync, mainApp, settingsRepository)
		settingsWindow.Show()
		settingsWindow.RequestFocus()
	})
	settingsButton.Resize(fyne.NewSize(30, 30))
	return settingsButton
}

func NewPauseButton(pause chan bool, inSync chan string) *widget.Button {
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

	go func(syncAddress chan string) {
		for {
			if <-syncAddress != "" {
				pauseFocusButton.Disable()
			} else {
				pauseFocusButton.Enable()
			}
		}
	}(inSync)
	pauseFocusButton.SetIcon(icons[toPause])
	pauseFocusButton.Resize(fyne.NewSize(30, 30))
	return pauseFocusButton
}
