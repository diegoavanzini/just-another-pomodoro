package container

import (
	"fyne.io/fyne"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

type ButtonContainer struct {
	Container *fyne.Container
}

func NewButtonContainer(pause chan bool) ButtonContainer {
	var pauseFocusButton *widget.Button
	toPause := true
	pauseFocusButton = widget.NewButton("||", func() {
		pause <- toPause
		pause <- toPause
		toPause = !toPause
	})
	pauseFocusButton.Resize(fyne.NewSize(30, 30))

	var pomodoro fyne.App
	settingsButton := widget.NewButton("Settings...", func() {
		dialog.ShowInformation("Information", "TODO implementare setting del pomodoro", pomodoro.NewWindow("Settings"))
	})
	settingsButton.Resize(fyne.NewSize(30, 30))
	buttonsContainer := fyne.NewContainerWithLayout(
		layout.NewVBoxLayout(),
		pauseFocusButton,
		settingsButton)
	return ButtonContainer{Container: buttonsContainer}
}
