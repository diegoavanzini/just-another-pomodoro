package main

import (
	"bitbucket.org/avanz/anotherPomodoro/custom/container"
	"bitbucket.org/avanz/anotherPomodoro/repository"
	"bitbucket.org/avanz/anotherPomodoro/sync"
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

	//mainErrorListener := make(chan error)

	repository, err := repository.NewPomodoroRepository()
	if err != nil {
		panic(err)
	}

	synclistener := sync.NewListener(repository)
	synclistener.Start()

	syncRemoteAddressListener := make(chan string)
	repository.Write("settings", "synkAddress", "")
	if err != nil {
		panic(err)
	}

	logoBox := packr.New("logo", "."+string(os.PathSeparator)+"img"+string(os.PathSeparator)+"jap_logo.png")
	logo, err := fyne.LoadResourceFromPath(logoBox.ResolutionDir)

	pomodoroApp := app.New()
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

	mainWindowContainer := fyne.NewContainerWithLayout(layout.NewVBoxLayout())
	addPomodoro := make(chan bool)
	mainWindowContainer.AddObject(
		fyne.NewContainerWithLayout(
			layout.NewHBoxLayout(),
			fyne.NewContainerWithLayout(
				layout.NewVBoxLayout(),
				container.NewProgressBarContainer(pause, alert, addPomodoro, syncRemoteAddressListener, repository).Container,
				NewDailyPomodoroContainer(addPomodoro, repository)),
			container.NewButtonContainer(pause, syncRemoteAddressListener, pomodoroApp, repository)))

	pomodoroWindows.SetContent(mainWindowContainer)
	pomodoroWindows.ShowAndRun()
}

func NewDailyPomodoroContainer(addPomodoro chan bool, repository repository.IPomodoroRepository) *fyne.Container {
	doneContainer := container.NewPomodoroDoneContainer(layout.NewHBoxLayout(), repository)
	doneContainer.AddPomodoro()

	go func(pDone *container.PomodoroDoneContainer) {
		for {
			select {
			case <-addPomodoro:
				pDone.AddPomodoro()
			}
		}
	}(doneContainer)

	return fyne.NewContainerWithLayout(
		layout.NewVBoxLayout(),
		canvas.NewRectangle(color.White),
		doneContainer.Container,
		canvas.NewRectangle(color.White))
}
