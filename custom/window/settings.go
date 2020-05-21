package window

import (
	"bitbucket.org/avanz/anotherPomodoro/common"
	"bitbucket.org/avanz/anotherPomodoro/repository"
	"bitbucket.org/avanz/anotherPomodoro/sync"
	"errors"
	"fmt"
	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"github.com/atotto/clipboard"
	"net"
	"strings"
	"time"
)

func NewSettingsWindow(syncRemoteAddressListener chan string, pomodoroApp fyne.App, settingsRepository repository.IPomodoroRepository, synclistener sync.IListener) fyne.Window {
	settings := pomodoroApp.NewWindow("Settings")
	settings.Resize(fyne.Size{300, 50})

	timeDurationEntry := widget.NewEntry()
	timeDurationString := initTimeDuration("timeDuration", 25*time.Minute, settingsRepository)
	timeDurationEntry.SetText(timeDurationString)
	timeDurationInput := widget.NewFormItem("time duration", timeDurationEntry)

	timePauseEntry := widget.NewEntry()
	pauseDurationString := initTimeDuration("pauseDuration", 5*time.Minute, settingsRepository)
	timePauseEntry.SetText(pauseDurationString)
	timePauseInput := widget.NewFormItem("pause duration", timePauseEntry)

	timeShareAddressAndPort := ""
	settingsRepository.Read("settings", "timeShareAddressAndPort", &timeShareAddressAndPort)

	var err error
	IPAddressToShare := ""
	timeSharePortText := "1234"
	if timeShareAddressAndPort != "" {
		IPAddressToShareAndPort := strings.Split(timeShareAddressAndPort, ":")
		IPAddressToShare = IPAddressToShareAndPort[0]
		timeSharePortText = IPAddressToShareAndPort[1]
	}
	IPAddressToShare, err = common.ExternalIP()
	if err != nil {
		common.MainErrorListener <- err
	}
	timeShareLabel := widget.NewEntry()
	timeShareLabel.SetText(IPAddressToShare)
	timeShareLabel.Disable()

	timeSharePort := widget.NewEntry()
	timeSharePort.Text = timeSharePortText
	timeShareCopy := widget.NewButton("Copy", func() {
		timeShareText := fmt.Sprintf("%s:%s", timeShareLabel.Text, timeSharePort.Text)
		clipboard.WriteAll(timeShareText)
	})

	timeShareContainer := fyne.NewContainerWithLayout(layout.NewHBoxLayout(), timeShareLabel, timeSharePort, timeShareCopy)
	timeShareInput := widget.NewFormItem("share with other", timeShareContainer)

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

	saveButton := createSaveButton(timeDurationEntry, timePauseEntry, sinkEntry, timeShareLabel, timeSharePort, settings, settingsRepository, syncRemoteAddressListener, synclistener)
	insertOk := widget.NewFormItem("", saveButton)

	settingsForm := widget.NewForm(timeDurationInput, timePauseInput, timeShareInput, synkInput, insertOk)
	settingsContainer := fyne.NewContainerWithLayout(layout.NewVBoxLayout(), settingsForm)
	settings.SetContent(settingsContainer)
	return settings
}

func createSaveButton(timeDurationEntry, timePauseEntry, synkAddress, timeShareLabel, timeSharePort *widget.Entry, settings fyne.Window, settingsRepository repository.IPomodoroRepository, syncRemoteAddressListener chan string, synclistener sync.IListener) *widget.Button {
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

		timeShareAddressAndPort := fmt.Sprintf("%s:%s", timeShareLabel.Text, timeSharePort.Text)
		settingsRepository.Write("settings", "timeShareAddressAndPort", timeShareAddressAndPort)
		common.MainInfoListener <- "this changes will be effective restarting pomodoro."
		//timeSharePortInt, err := strconv.Atoi(timeSharePort.Text)
		//if err != nil {
		//	common.MainErrorListener <- err
		//}
		//synclistener.ChangeAddressAndPort(timeShareLabel.Text, timeSharePortInt)

		sa := strings.Split(synkAddress.Text, ":")
		if sa[0] != "" && net.ParseIP(sa[0]) == nil {
			common.MainErrorListener <- errors.New("invalid remote address:" + sa[0])
			synkAddress.Text = ""
		} else {
			settingsRepository.Write("settings", "synkAddress", synkAddress.Text)
			syncRemoteAddressListener <- synkAddress.Text
		}
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
