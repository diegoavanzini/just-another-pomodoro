package main

import (
	"bitbucket.org/avanz/anotherPomodoro/common"
	"bitbucket.org/avanz/anotherPomodoro/custom/container"
	"bitbucket.org/avanz/anotherPomodoro/repository"
	"bitbucket.org/avanz/anotherPomodoro/sync"
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	packr "github.com/gobuffalo/packr/v2"
	"image/color"
	"os"
)

const title = "just another pomodoro"

func main() {

	logoBox := packr.New("logo", "."+string(os.PathSeparator)+"img"+string(os.PathSeparator)+"jap_logo.png")
	logo, _ := fyne.LoadResourceFromPath(logoBox.ResolutionDir)

	pomodoroApp := app.New()
	pomodoroApp.SetIcon(logo)
	pomodoroWindows := pomodoroApp.NewWindow(title)
	pomodoroWindows.SetFixedSize(true)
	pomodoroWindows.SetMaster()

	initInformationWindowListener(pomodoroApp)
	initErrorWindowListener(pomodoroApp)

	repository, err := repository.NewPomodoroRepository()
	if err != nil {
		common.MainErrorListener <- err
	}

	synclistener := sync.NewListener(repository)
	synclistener.Start()

	syncRemoteAddressListener := make(chan string)
	//if err != nil {
	//	common.MainErrorListener <- repository.Write("settings", "synkAddress", "")
	//}
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
	addPomodoroListener := make(chan bool)
	mainWindowContainer.AddObject(
		fyne.NewContainerWithLayout(
			layout.NewHBoxLayout(),
			fyne.NewContainerWithLayout(
				layout.NewVBoxLayout(),
				container.NewProgressBarContainer(pause, alert, addPomodoroListener, syncRemoteAddressListener, repository).Container,
				NewDailyPomodoroContainer(addPomodoroListener, repository)),
			container.NewButtonContainer(pause, syncRemoteAddressListener, pomodoroApp, repository, synclistener)))

	pomodoroWindows.SetContent(mainWindowContainer)
	pomodoroWindows.ShowAndRun()
}

func initInformationWindowListener(pomodoroApp fyne.App) {
	common.MainInfoListener = make(chan string)
	go func(mainInformationListener chan string, pomodoroApp fyne.App) {
		for {
			message := <-mainInformationListener
			if message != "" {
				informationWindows := pomodoroApp.NewWindow("Information")
				informationWindows.Resize(fyne.NewSize(300, 50))
				informationWindows.SetContent(fyne.NewContainerWithLayout(layout.NewVBoxLayout(), widget.NewLabel(message), widget.NewButton("close", func() {
					informationWindows.Close()
				})))
				informationWindows.Show()
				informationWindows.SetFixedSize(true)
				informationWindows.RequestFocus()
			}
		}
	}(common.MainInfoListener, pomodoroApp)
}

func initErrorWindowListener(pomodoroApp fyne.App) {
	common.MainErrorListener = make(chan error)
	go func(mainErrorListener chan error, pomodoroApp fyne.App) {
		for {
			err := <-mainErrorListener
			if err != nil {
				errorWindows := pomodoroApp.NewWindow("Error")
				errorWindows.Resize(fyne.NewSize(300, 50))
				errorWindows.SetContent(fyne.NewContainerWithLayout(layout.NewVBoxLayout(), widget.NewLabel(err.Error()), widget.NewButton("close", func() {
					errorWindows.Close()
				})))
				errorWindows.Show()
				errorWindows.SetFixedSize(true)
				//errorWindows.CenterOnScreen()
				errorWindows.RequestFocus()
			}
		}
	}(common.MainErrorListener, pomodoroApp)
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
