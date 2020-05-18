package window

import (
	"bitbucket.org/avanz/anotherPomodoro/common"
	"bitbucket.org/avanz/anotherPomodoro/repository"
	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"log"
	"time"
)

func NewSettingsWindow(inSync chan string, pomodoroApp fyne.App, settingsRepository repository.IPomodoroRepository) fyne.Window {
	settings := pomodoroApp.NewWindow("Settings")
	settings.Resize(fyne.Size{300, 50})
	settings.CenterOnScreen()

	timeDurationEntry := widget.NewEntry()
	timeDurationString := initTimeDuration("timeDuration", 25*time.Minute, settingsRepository)
	timeDurationEntry.SetText(timeDurationString)
	timeDurationInput := widget.NewFormItem("time duration", timeDurationEntry)

	timePauseEntry := widget.NewEntry()
	pauseDurationString := initTimeDuration("pauseDuration", 5*time.Minute, settingsRepository)
	timePauseEntry.SetText(pauseDurationString)
	timePauseInput := widget.NewFormItem("pause duration", timePauseEntry)

	IPAddressToShare, err := common.ExternalIP()
	if err != nil {
		log.Fatal(err)
	}
	timeShareButton := widget.NewLabel(IPAddressToShare)
	timeShareInput := widget.NewFormItem("share with other", timeShareButton)

	sinkEntry := widget.NewEntry()
	synkAddress := ""
	err = settingsRepository.Read("settings", "synkAddress", &synkAddress)
	if synkAddress == "" {
		sinkEntry.SetText("")
		sinkEntry.SetPlaceHolder("paste here the remote address")
	} else {
		sinkEntry.SetText(synkAddress)
	}
	synkInput := widget.NewFormItem("synk with", sinkEntry)

	insertOk := widget.NewFormItem("", createSaveButton(timeDurationEntry, timePauseEntry, sinkEntry, settings, settingsRepository, inSync))

	settingsForm := widget.NewForm(timeDurationInput, timePauseInput, timeShareInput, synkInput, insertOk)
	settingsContainer := fyne.NewContainerWithLayout(layout.NewVBoxLayout(), settingsForm)
	settings.SetContent(settingsContainer)
	return settings
}

func createSaveButton(timeDurationEntry, timePauseEntry, synkAddress *widget.Entry, settings fyne.Window, settingsRepository repository.IPomodoroRepository, inSync chan string) *widget.Button {
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
		settingsRepository.Write("settings", "synkAddress", synkAddress.Text)
		inSync <- synkAddress.Text
		settings.Close()
	})
}

func initTimeDuration(resource string, defaultDuration time.Duration, settingsRepository repository.IPomodoroRepository) string {
	var timeDurationEntryStored time.Duration
	err := settingsRepository.Read("settings", resource, &timeDurationEntryStored)
	if err != nil {
		panic(err)
	}
	if timeDurationEntryStored == 0*time.Minute {
		timeDurationEntryStored = defaultDuration
	}
	return common.DurationToString(timeDurationEntryStored)
}
