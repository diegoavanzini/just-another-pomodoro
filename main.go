package main

import (
	"bitbucket.org/avanz/anotherPomodoro/common"
	"bitbucket.org/avanz/anotherPomodoro/custom/container"
	"bitbucket.org/avanz/anotherPomodoro/repository"
	"bitbucket.org/avanz/anotherPomodoro/sync"
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/layout"
	"github.com/gobuffalo/packr/v2"
	"image/color"
	"log"
	"os"
)

const title = "just another pomodoro"

func main() {

	logoBox := packr.New("logo", "."+string(os.PathSeparator)+"img"+string(os.PathSeparator)+"jap_logo.png")
	logo, err := fyne.LoadResourceFromPath(logoBox.ResolutionDir)

	pomodoroApp := app.New()
	if err != nil {
		log.Fatal(err)
	}
	pomodoroApp.SetIcon(logo)
	pomodoroWindows := pomodoroApp.NewWindow(title)
	common.MainErrorListener = make(chan error)
	go func (mainErrorListener chan error, pomodoroApp fyne.App) {
		for  {
			err := <- mainErrorListener
			if err != nil {
				win := pomodoroApp.NewWindow("Dialogs")
				dialog.ShowError(err, win)
			}
		}
	} (common.MainErrorListener, pomodoroApp)


	repository, err := repository.NewPomodoroRepository()
	common.MainErrorListener <- err

	synclistener := sync.NewListener(repository)
	synclistener.Start()

	syncRemoteAddressListener := make(chan string)
	common.MainErrorListener <- repository.Write("settings", "synkAddress", "")

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
