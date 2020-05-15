package main

import (
	"bitbucket.org/avanz/anotherPomodoro/common"
	"bitbucket.org/avanz/anotherPomodoro/custom/container"
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/layout"
	"github.com/gobuffalo/packr/v2"
	"github.com/nanobox-io/golang-scribble"
	"image/color"
	"os"
)

const title = "just another pomodoro"

func main() {
	initStorage()

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
	buttonsContainer := container.NewButtonContainer(pause, pomodoroApp)

	addPomodoro := make(chan bool)
	progressBarContainer := container.NewProgressBarContainer(pause, alert, addPomodoro)

	dailyContainer := NewDailyPomodoroContainer(addPomodoro)

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

func initStorage() {
	dataFolder := packr.New("data", "./data")
	var err error
	common.Settings, err = scribble.New(dataFolder.ResolutionDir, nil)
	if err != nil {
		panic(err)
	}
}

func NewDailyPomodoroContainer(addPomodoro chan bool) *fyne.Container {
	doneContainer := container.NewPomodoroDoneContainer(layout.NewHBoxLayout())
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
