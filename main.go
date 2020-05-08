package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"image/color"
	"time"
)

var pomodoroWindows fyne.Window
var pomodoro fyne.App

const title = "just another pomodoro"

func main() {
	pomodoro = app.New()
	pomodoroWindows = pomodoro.NewWindow(title)
	pomodoroWindows.Resize(fyne.Size{420, 80})
	winGrid := fyne.NewContainerWithLayout(layout.NewVBoxLayout())

	pause := make(chan bool)

	buttonsContainer := createButtonContainer(pause)

	progressBar, pauseProgressBar, progressBarContainer := createProgressBarContainer()

	doneContainer := fyne.NewContainerWithLayout(NewResizableGridLayout(21))
	doneContainer.Resize(fyne.Size{20, 300})
	progressAndDoneContainer := fyne.NewContainerWithLayout(layout.NewVBoxLayout(), progressBarContainer, doneContainer)
	progressAndDoneContainer.Resize(fyne.Size{310, 50})

	timerContainer := fyne.NewContainerWithLayout(NewResizableGridLayout(2), progressAndDoneContainer, buttonsContainer)
	winGrid.AddObject(timerContainer)

	go func(pDone *fyne.Container) {
		for {
			startTimerCycle(progressBar, pauseProgressBar, pause)
			AddPomodoro(pDone)
		}
	}(doneContainer)

	pomodoroWindows.SetContent(winGrid)
	pomodoroWindows.ShowAndRun()
}

func createProgressBarContainer() (*CustomProgressBar, *CustomProgressBar, *fyne.Container) {
	progressBar := NewTimerProgressBar(25 * time.Second)
	progressBar.Resize(fyne.NewSize(250, 5))
	pauseProgressBar := NewTimerProgressBar(5 * time.Second)
	pauseProgressBar.Resize(fyne.NewSize(50, 5))
	progressBarContainer := fyne.NewContainerWithLayout(NewResizableGridLayout(2), progressBar, pauseProgressBar)
	return progressBar, pauseProgressBar, progressBarContainer
}

func createButtonContainer(pause chan bool) *fyne.Container {
	var pauseFocusButton *widget.Button
	pauseFocusButton = widget.NewButton("||", func() {
		pomodoroWindows.SetTitle(title + " (paused)")
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
	circle := canvas.NewCircle(color.RGBA{R: 11, G: 156, B: 49, A: 1})
	circle.Resize(fyne.Size{Width: 10, Height: 10})
	circle.Show()
	if len(c.Objects) < 20 {
		c.AddObject(circle)
		return
	}
	newList := c.Objects[1 : len(c.Objects)-1]
	c.Objects = newList
	c.AddObject(circle)
	label := widget.NewLabel("...")
	c.AddObject(label)
}

func startTimerCycle(progressBar *CustomProgressBar, pauseProgressBar *CustomProgressBar, pause chan bool) {
	startTimer(progressBar, pause)
	startTimer(pauseProgressBar, pause)
	pauseProgressBar.SetValue(time.Duration(0))
}

func startTimer(bar *CustomProgressBar, pause chan bool) {
	ticker := time.NewTicker(1 * time.Second)
	value := time.Duration(0)
	func() {
		for {
			select {
			case <-pause:
				<-pause
				pomodoroWindows.SetTitle(title)
			case <-ticker.C:
				value += 1*time.Second
				bar.SetValue(value)
				if value >= bar.Max {
					ticker.Stop()
					return
				}
			}
		}
	}()
}
