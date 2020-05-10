package main

import (
	"bitbucket.org/avanz/anotherPomodoro/custom/container"
	custom_layout "bitbucket.org/avanz/anotherPomodoro/custom/layout"
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/layout"
	"image/color"
)


const title = "just another pomodoro"

func main() {
	pomodoro := app.New()
	pomodoroWindows := pomodoro.NewWindow(title)

	alert := make(chan bool)
	go func(alert chan bool) {
		for {
			<-alert
			pomodoroWindows.RequestFocus()
		}
	}(alert)
	pomodoroWindows.Resize(fyne.Size{420, 70})
	winGrid := fyne.NewContainerWithLayout(layout.NewVBoxLayout())

	pause := make(chan bool, 2)
	addPomodoro := make(chan bool)
	buttonsContainer := container.NewButtonContainer(pause)
	progressBarContainer := container.NewProgressBarContainer(pause, alert, addPomodoro)

	doneContainer := container.NewPomodoroDoneContainer(layout.NewHBoxLayout())
	doneContainer.AddPomodoro()
	dailyContainer := fyne.NewContainerWithLayout(
		layout.NewVBoxLayout(),
		canvas.NewRectangle(color.White),
		doneContainer.Container)

	progressAndDoneContainer := fyne.NewContainerWithLayout(
		layout.NewVBoxLayout(),
		progressBarContainer.Container,
		dailyContainer)
	progressAndDoneContainer.Resize(fyne.Size{310, 50})

	timerContainer := fyne.NewContainerWithLayout(
		custom_layout.NewResizableGridLayout(2),
		progressAndDoneContainer,
		buttonsContainer.Container)

	go func(pDone *container.PomodoroDoneContainer) {
		for {
			select {
			case <-addPomodoro:
				pDone.AddPomodoro()
			case paused :=<-pause:
				if paused {
					pomodoroWindows.SetTitle(title + " (paused)")
				} else {
					pomodoroWindows.SetTitle(title)
				}
			}
		}
	}(doneContainer)
	winGrid.AddObject(timerContainer)

	pomodoroWindows.SetContent(winGrid)
	pomodoroWindows.ShowAndRun()
}
