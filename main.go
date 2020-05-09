package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"image/color"
	"math"
	"time"
)

var pomodoroWindows fyne.Window
var pomodoro fyne.App

const title = "just another pomodoro"

func main() {
	pomodoro = app.New()
	pomodoroWindows = pomodoro.NewWindow(title)

	alert := make(chan bool)
	go func(alert chan bool) {
		for {
			<-alert
			pomodoroWindows.RequestFocus()
		}
	}(alert)
	pomodoroWindows.Resize(fyne.Size{420, 70})
	winGrid := fyne.NewContainerWithLayout(layout.NewVBoxLayout())

	pause := make(chan bool)
	buttonsContainer := createButtonContainer(pause)
	progressBar, pauseProgressBar, progressBarContainer := createProgressBarContainer(pause, alert)

	doneContainer := fyne.NewContainerWithLayout(layout.NewHBoxLayout())
	AddPomodoro(doneContainer)
	dailyContainer := fyne.NewContainerWithLayout(
		layout.NewVBoxLayout(),
		canvas.NewRectangle(color.White),
		doneContainer)

	progressAndDoneContainer := fyne.NewContainerWithLayout(
		layout.NewVBoxLayout(),
		progressBarContainer,
		dailyContainer)
	progressAndDoneContainer.Resize(fyne.Size{310, 50})

	timerContainer := fyne.NewContainerWithLayout(
		NewResizableGridLayout(2),
		progressAndDoneContainer,
		buttonsContainer)

	go func(pDone *fyne.Container) {
		for {
			startTimerCycle(progressBar, pauseProgressBar, pause, alert)
			AddPomodoro(pDone)
		}
	}(doneContainer)
	winGrid.AddObject(timerContainer)

	pomodoroWindows.SetContent(winGrid)
	pomodoroWindows.ShowAndRun()
}

func createProgressBarContainer(pause, alert chan bool) (*CustomProgressBar, *CustomProgressBar, *fyne.Container) {
	progressBar := NewTimerProgressBar(25*time.Minute, pause, alert)
	progressBar.Resize(fyne.NewSize(250, 5))
	pauseProgressBar := NewTimerProgressBar(5*time.Minute, pause, alert)
	pauseProgressBar.Resize(fyne.NewSize(50, 5))
	progressBarContainer := fyne.NewContainerWithLayout(NewResizableGridLayout(2), progressBar, pauseProgressBar)
	return progressBar, pauseProgressBar, progressBarContainer
}

func createButtonContainer(pause chan bool) *fyne.Container {
	var pauseFocusButton *widget.Button
	pauseFocusButton = widget.NewButton("||", func() {
		if pomodoroWindows.Title() == title {
			pomodoroWindows.SetTitle(title + " (paused)")
		} else {
			pomodoroWindows.SetTitle(title)
		}
		pause <- true
	})
	pauseFocusButton.Resize(fyne.NewSize(30, 30))

	settingsButton := widget.NewButton("Settings...", func() {
		dialog.ShowInformation("Information", "TODO implementare setting del pomodoro", pomodoro.NewWindow("Settings"))
	})
	settingsButton.Resize(fyne.NewSize(30, 30))
	buttonsLayout := layout.NewVBoxLayout()
	buttonsContainer := fyne.NewContainerWithLayout(buttonsLayout, pauseFocusButton, settingsButton)
	return buttonsContainer
}

func AddPomodoro(c *fyne.Container) {
	totalMinutes := time.Now().Hour() * 60 + time.Now().Minute()
	numberOfPresentPomodoro := int(math.Floor(float64(totalMinutes / (25 + 5))))
	if len(c.Objects) == numberOfPresentPomodoro + 1 {
		return
	}
	if len(c.Objects) == 0 {
		for i := 0; i < numberOfPresentPomodoro; i++ {
			pomodoro := createPomodoro(5)
			pomodoro.Hide()
			c.AddObject(pomodoro)
		}
	}
	pomodoro := createPomodoro(5)
	pomodoro.Show()
	c.AddObject(pomodoro)
}

func createPomodoro(size int) *canvas.Rectangle {
	pomodoro := canvas.NewRectangle(color.RGBA{R: 11, G: 156, B: 49, A: 1})
	pomodoro.SetMinSize(fyne.Size{Width: size, Height: 20})
	return pomodoro
}

func startTimerCycle(progressBar *CustomProgressBar, pauseProgressBar *CustomProgressBar, pause, alert chan bool) {
	progressBar.Start()
	pauseProgressBar.Start()
	pauseProgressBar.SetValue(time.Duration(0))
}
