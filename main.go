package main

import (
	"bitbucket.org/avanz/anotherPomodoro/custom/container"
	"bitbucket.org/avanz/anotherPomodoro/repository"
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/layout"
	"github.com/gobuffalo/packr/v2"
	"image/color"
	"os"
)

const title = "just another pomodoro"

func main() {
	repository, err := repository.NewSettingsRepository()
	if err != nil {
		panic(err)
	}

	logoBox := packr.New("logo", "."+string(os.PathSeparator)+"img"+string(os.PathSeparator)+"jap_logo.png")
	pomodoroApp := app.New()
	logo, err := fyne.LoadResourceFromPath(logoBox.ResolutionDir)
	if err != nil {
		panic(err)
	}
	pomodoroApp.SetIcon(logo)
	pomodoroWindows := pomodoroApp.NewWindow(title)

	pause := make(chan bool, 2)
	alert := make(chan bool)
	go func(pause, alert chan bool) {
		for {
			select {
			case <-alert:
				pomodoroWindows.RequestFocus()
			case paused := <-pause:
				if paused {
					pomodoroWindows.SetTitle(title + " (paused)")
				} else {
					pomodoroWindows.SetTitle(title)
				}
			}
		}
	}(pause, alert)

	globalContainer := fyne.NewContainerWithLayout(layout.NewVBoxLayout())
	buttonsContainer := container.NewButtonContainer(pause, pomodoroApp,repository)

	addPomodoro := make(chan bool)
	progressBarContainer := container.NewProgressBarContainer(pause, alert, addPomodoro,repository)

	dailyContainer := NewDailyPomodoroContainer(addPomodoro, repository)

	progressAndDoneContainer := fyne.NewContainerWithLayout(
		layout.NewVBoxLayout(),
		progressBarContainer.Container,
		dailyContainer)
	//progressAndDoneContainer.Resize(fyne.Size{310, 50})

	timerContainer := fyne.NewContainerWithLayout(
		layout.NewHBoxLayout(),
		progressAndDoneContainer,
		buttonsContainer)

	globalContainer.AddObject(timerContainer)

	pomodoroWindows.SetContent(globalContainer)
	pomodoroWindows.ShowAndRun()
}

func initSettings() (repository.ISettingsRepository, error) {
	return repository.NewSettingsRepository()
}

func NewDailyPomodoroContainer(addPomodoro chan bool, repository repository.ISettingsRepository) *fyne.Container {
	doneContainer := container.NewPomodoroDoneContainer(layout.NewHBoxLayout(), repository)
	doneContainer.AddPomodoro()

	dailyContainer := fyne.NewContainerWithLayout(
		layout.NewVBoxLayout(),
		canvas.NewRectangle(color.White),
		doneContainer.Container,
		canvas.NewRectangle(color.White))

	go func(pDone *container.PomodoroDoneContainer) {
		for {
			select {
			case <-addPomodoro:
				pDone.AddPomodoro()
			}
		}
	}(doneContainer)

	return dailyContainer
}
