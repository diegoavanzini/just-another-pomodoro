package container

import (
	"bitbucket.org/avanz/anotherPomodoro/common"
	"bitbucket.org/avanz/anotherPomodoro/repository"
	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"github.com/gobuffalo/packr/v2"

	"os"
	"time"
)

type ButtonContainer struct {
	Container *fyne.Container
}

func NewButtonContainer(pause chan bool, mainApp fyne.App, settingsRepository repository.ISettingsRepository) *fyne.Container {
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
	icons := map[bool]fyne.Resource{false:startIcon,true:pauseIcon}
	var pauseFocusButton *widget.Button
	toPause := true
	pauseFocusButton = widget.NewButton("", func() {
		pause <- toPause
		pause <- toPause
		toPause = !toPause
		pauseFocusButton.SetIcon(icons[toPause])
	})
	pauseFocusButton.SetIcon(icons[toPause])
	pauseFocusButton.Resize(fyne.NewSize(30, 30))

	settingsButton := widget.NewButton("Settings...", func() {
		settings := createSettingsWindows(mainApp, settingsRepository)
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

func createSettingsWindows(pomodoroApp fyne.App, settingsRepository repository.ISettingsRepository) fyne.Window {
	settings := pomodoroApp.NewWindow("Settings")
	settings.Resize(fyne.Size{300, 50})
	settings.CenterOnScreen()

	timeDurationEntry := widget.NewEntry()

	initTimeDuration(timeDurationEntry, "timeDuration", 25*time.Minute, settingsRepository)
	timeDurationInput := widget.NewFormItem("time duration:", timeDurationEntry)

	timePauseEntry := widget.NewEntry()
	initTimeDuration(timePauseEntry, "pauseDuration", 5*time.Minute, settingsRepository)
	timePauseInput := widget.NewFormItem("pause duration:", timePauseEntry)

	insertOk := widget.NewFormItem("", createSaveButton(timeDurationEntry, timePauseEntry, settings, settingsRepository))

	settingsForm := widget.NewForm(timeDurationInput, timePauseInput, insertOk)
	settingsContainer := fyne.NewContainerWithLayout(layout.NewVBoxLayout(), settingsForm)
	settings.SetContent(settingsContainer)
	return settings
}

func createSaveButton(timeDurationEntry *widget.Entry, timePauseEntry *widget.Entry, settings fyne.Window, settingsRepository repository.ISettingsRepository) *widget.Button {
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
		settings.Close()
	})
}

func initTimeDuration(timeDurationEntry *widget.Entry, resource string, defaultDuration time.Duration, settingsRepository repository.ISettingsRepository) {
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
