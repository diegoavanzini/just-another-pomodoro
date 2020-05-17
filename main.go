package main

import (
	"bitbucket.org/avanz/anotherPomodoro/custom/container"
	"bitbucket.org/avanz/anotherPomodoro/repository"
	"bufio"
	"encoding/json"
	"fmt"
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/layout"
	"github.com/gobuffalo/packr/v2"
	"image/color"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

const title = "just another pomodoro"
const SERVER_PORT = "1234"

func main() {

	repository, err := repository.NewPomodoroRepository()
	if err != nil {
		panic(err)
	}

	startSharingListener(repository)

	repository.Write("settings", "synkAddressPort", "")
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
	inSync := make(chan bool)
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
	buttonsContainer := container.NewButtonContainer(pause, inSync, pomodoroApp, repository)

	addPomodoro := make(chan bool)
	progressBarContainer := container.NewProgressBarContainer(pause, alert, addPomodoro, repository)

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

func startSharingListener(repository repository.IPomodoroRepository) {
	countConnection := 0
	go func() {
		l, err := net.Listen("tcp4", fmt.Sprintf("%s:%s", "127.0.0.1", SERVER_PORT))
		if err != nil {
			log.Fatal(err)
			//return
		}
		defer l.Close()

		for {
			c, err := l.Accept()
			if err != nil {
				log.Fatal(err)
				return
			}
			go handleConnection(c, repository)
			countConnection++
		}
	}()
}

func handleConnection(c net.Conn, repository repository.IPomodoroRepository) {
	fmt.Print(".")
	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		temp := strings.TrimSpace(string(netData))
		if temp == "STOP" {
			break
		}
		var currentTimerValue string
		err = repository.Read("current", "timerValue", &currentTimerValue)
		if err != nil {
			log.Fatal(err)
		}
		var timeDuration = int((25 * time.Minute).Seconds())
		err = repository.Read("settings", "timeDuration", &timeDuration)
		if err != nil {
			log.Fatal(err)
		}
		var pauseDuration = int((5 * time.Minute).Seconds())
		err = repository.Read("settings", "pauseDuration", &pauseDuration)
		if err != nil {
			log.Fatal(err)
		}
		currentPomodoro := struct {
			TimeDuration      int
			PauseDuration     int
			CurrentTimerValue string
		}{timeDuration, pauseDuration, currentTimerValue}
		currentPomodoroJson, err := json.Marshal(currentPomodoro)
		if err != nil {
			log.Fatal(err)
		}
		c.Write([]byte(string(currentPomodoroJson) + "\n"))
	}
	c.Close()
}

func NewDailyPomodoroContainer(addPomodoro chan bool, repository repository.IPomodoroRepository) *fyne.Container {
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
