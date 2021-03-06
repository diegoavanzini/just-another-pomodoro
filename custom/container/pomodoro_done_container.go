package container

import (
	"bitbucket.org/avanz/anotherPomodoro/common"
	"bitbucket.org/avanz/anotherPomodoro/custom/widget"
	"bitbucket.org/avanz/anotherPomodoro/repository"
	"encoding/json"
	"fmt"
	"fyne.io/fyne"
	"fyne.io/fyne/theme"
	"time"
)

type PomodoroDoneContainer struct {
	*fyne.Container
	repository repository.IPomodoroRepository
}

func NewPomodoroDoneContainer(boxLayout fyne.Layout, repository repository.IPomodoroRepository) *PomodoroDoneContainer {
	container := &PomodoroDoneContainer{
		Container:  fyne.NewContainerWithLayout(boxLayout),
		repository: repository,
	}
	for i := 0; i < 48; i++ {
		pomodoro := widget.NewPomodoro(5, theme.BackgroundColor())
		container.AddObject(pomodoro)
	}
	return container
}

type PomodoroStruct struct {
	TimeStarted  string
	TimeDuration int
}

func (c PomodoroDoneContainer) AddPomodoro() {
	timeDuration := 25 * time.Minute
	err := c.repository.Read("settings", "timeDuration", &timeDuration)
	if err != nil {
		panic(err)
	}
	pauseDuration := 5 * time.Minute
	err = c.repository.Read("settings", "pauseDuration", &pauseDuration)
	if err != nil {
		panic(err)
	}

	workdoneToday := "workdone" + time.Now().Format("20060102")
	pomodoroList, _ := c.repository.ReadAll(workdoneToday)
	layout := "02-01-2006 15:04"
	started := time.Now().UTC()
	currentPosition := c.getPosition(started)
	skipInsert := false
	for _, pString := range pomodoroList {
		p := PomodoroStruct{}
		if err := json.Unmarshal([]byte(pString), &p); err != nil {
			common.MainErrorListener <- err
		}
		started, err := time.Parse(layout, p.TimeStarted)
		if err != nil {
			common.MainErrorListener <- err
		}
		pPosition := c.getPosition(started)
		if pPosition == currentPosition && !skipInsert {
			skipInsert = true
		}
		c.Objects[pPosition] = widget.NewPomodoro(5, common.Green)
	}
	if !skipInsert {
		pomodoroStruct := PomodoroStruct{
			TimeStarted:  started.Format(layout),
			TimeDuration: int(timeDuration.Minutes()),
		}
		if err := c.repository.Write(workdoneToday, fmt.Sprintf("pomodoro_%d", currentPosition), pomodoroStruct); err != nil {
			common.MainErrorListener <- err
		}
	}
}

func (c PomodoroDoneContainer) getPosition(started time.Time) int {
	pPosition := started.Hour() * 2
	if started.Minute() > 30 {
		pPosition = pPosition + 1
	}
	return pPosition
}
