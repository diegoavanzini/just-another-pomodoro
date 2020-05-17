package container

import (
	"bitbucket.org/avanz/anotherPomodoro/common"
	"bitbucket.org/avanz/anotherPomodoro/repository"
	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"github.com/atotto/clipboard"
	"github.com/gobuffalo/packr/v2"
	"log"
	"os"
	"time"
)

type ButtonContainer struct {
	Container *fyne.Container
}

func NewButtonContainer(pause, inSync chan bool, mainApp fyne.App, settingsRepository repository.IPomodoroRepository) *fyne.Container {
	var err error
	startBox := packr.New("start", "."+string(os.PathSeparator)+"img"+string(os.PathSeparator)+"start.png")
	startIcon, err := fyne.LoadResourceFromPath(startBox.ResolutionDir)
	if err != nil {
		panic(err)
	}
	pauseBox := packr.New("pause", "."+string(os.PathSeparator)+"img"+string(os.PathSeparator)+"pause.png")
	pauseIcon, err := fyne.LoadResourceFromPath(pauseBox.ResolutionDir)
	if err != nil {
		panic(err)
	}
	icons := map[bool]fyne.Resource{false: startIcon, true: pauseIcon}
	var pauseFocusButton *widget.Button
	toPause := true
	pauseFocusButton = widget.NewButton("", func() {
		pause <- toPause
		pause <- toPause
		toPause = !toPause
		pauseFocusButton.SetIcon(icons[toPause])
	})

	go func(inSync chan bool) {
		for {
			sync := <-inSync
			if !sync {
				pauseFocusButton.Enable()
			} else {
				pauseFocusButton.Disable()
			}
		}
	}(inSync)
	pauseFocusButton.SetIcon(icons[toPause])
	pauseFocusButton.Resize(fyne.NewSize(30, 30))

	settingsButton := widget.NewButton("Settings...", func() {
		settings := createSettingsWindows(inSync, mainApp, settingsRepository)
		settings.Show()
		settings.RequestFocus()
	})
	settingsButton.Resize(fyne.NewSize(30, 30))
	buttonsContainer := fyne.NewContainerWithLayout(
		layout.NewVBoxLayout(),
		pauseFocusButton,
		settingsButton)
	return buttonsContainer
}

func createSettingsWindows(inSync chan bool, pomodoroApp fyne.App, settingsRepository repository.IPomodoroRepository) fyne.Window {
	settings := pomodoroApp.NewWindow("Settings")
	settings.Resize(fyne.Size{300, 50})
	settings.CenterOnScreen()

	timeDurationEntry := widget.NewEntry()
	initTimeDuration(timeDurationEntry, "timeDuration", 25*time.Minute, settingsRepository)
	timeDurationInput := widget.NewFormItem("time duration", timeDurationEntry)

	timePauseEntry := widget.NewEntry()
	initTimeDuration(timePauseEntry, "pauseDuration", 5*time.Minute, settingsRepository)
	timePauseInput := widget.NewFormItem("pause duration", timePauseEntry)

	IPAddressToShare, err := common.ExternalIP()
	if err != nil {
		log.Fatal(err)
	}
	timeShareButton := widget.NewButton(IPAddressToShare, func() {
		clipboard.WriteAll(IPAddressToShare)
	})
	timeShareInput := widget.NewFormItem("share with other", timeShareButton)

	sinkEntry := widget.NewEntry()
	synkAddressPort := ""
	err = settingsRepository.Read("settings", "synkAddressPort", &synkAddressPort)
	if sinkEntry.Text == "" {
		sinkEntry.SetPlaceHolder(synkAddressPort)
	} else {
		sinkEntry.SetText(synkAddressPort)
	}
	synkInput := widget.NewFormItem("synk with", sinkEntry)

	insertOk := widget.NewFormItem("", createSaveButton(timeDurationEntry, timePauseEntry, sinkEntry, settings, settingsRepository, inSync))

	settingsForm := widget.NewForm(timeDurationInput, timePauseInput, timeShareInput, synkInput, insertOk)
	settingsContainer := fyne.NewContainerWithLayout(layout.NewVBoxLayout(), settingsForm)
	settings.SetContent(settingsContainer)
	return settings
}

func createSaveButton(timeDurationEntry, timePauseEntry, synkAddressPort *widget.Entry, settings fyne.Window, settingsRepository repository.IPomodoroRepository, inSync chan bool) *widget.Button {
	return widget.NewButton("save", func() {
		duration, err := common.StringToDuration(timeDurationEntry.Text)
		if err != nil {
			panic(err)
		}
		settingsRepository.Write("settings", "timeDuration", duration)
		pauseDuration, err := common.StringToDuration(timePauseEntry.Text)
		if err != nil {
			panic(err)
		}
		settingsRepository.Write("settings", "pauseDuration", pauseDuration)
		settingsRepository.Write("settings", "synkAddressPort", synkAddressPort.Text)
		if synkAddressPort.Text != "" {
			inSync <- true
		} else {
			inSync <- false
		}
		settings.Close()
	})
}

func initTimeDuration(timeDurationEntry *widget.Entry, resource string, defaultDuration time.Duration, settingsRepository repository.IPomodoroRepository) {
	var timeDurationEntryStored time.Duration
	err := settingsRepository.Read("settings", resource, &timeDurationEntryStored)
	if err != nil {
		panic(err)
	}
	if timeDurationEntryStored == 0*time.Minute {
		timeDurationEntryStored = defaultDuration
	}
	timeDurationEntry.SetText(common.DurationToString(timeDurationEntryStored))
}
