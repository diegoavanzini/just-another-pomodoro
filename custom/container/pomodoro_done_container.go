package container

import (
	"bitbucket.org/avanz/anotherPomodoro/common"
	"bitbucket.org/avanz/anotherPomodoro/custom/widget"
	"encoding/json"
	"fmt"
	"fyne.io/fyne"
	"fyne.io/fyne/theme"
	"log"
	"time"
)

type PomodoroDoneContainer struct {
	*fyne.Container
}

func NewPomodoroDoneContainer(boxLayout fyne.Layout) *PomodoroDoneContainer {
	container := &PomodoroDoneContainer{
		Container: fyne.NewContainerWithLayout(boxLayout),
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
	err := common.Settings.Read("settings", "timeDuration", &timeDuration)
	if err != nil {
		panic(err)
	}
	pauseDuration := 5 * time.Minute
	err = common.Settings.Read("settings", "pauseDuration", &pauseDuration)
	if err != nil {
		panic(err)
	}

	pomodoroList, err := common.Settings.ReadAll("workdone")
	if err != nil {
		log.Fatal(err)
	}
	layout := "02-01-2006 15:04"
	started := time.Now().UTC()
	currentPosition := c.getPosition(started)
	skipInsert := false
	for _, pString := range pomodoroList {
		p := PomodoroStruct{}
		if err := json.Unmarshal([]byte(pString), &p); err != nil {
			log.Fatal(err)
		}
		started, err := time.Parse(layout, p.TimeStarted)
		if err!= nil {
			log.Fatal(err)
		}
		pPosition := c.getPosition(started)
		if pPosition == currentPosition && !skipInsert {
			skipInsert = true
		}
		c.Objects[pPosition] = widget.NewPomodoro(5, common.Green)
	}
	if !skipInsert {
		if err := common.Settings.Write("workdone", fmt.Sprintf("pomodoro_%s_%d", started.Format("20060102"), currentPosition), PomodoroStruct{
			TimeStarted:  started.Format(layout),
			TimeDuration: int(timeDuration.Minutes()),
		}); err != nil  {
			log.Fatal(err)
		}
	}
}

func (c PomodoroDoneContainer) getPosition(started time.Time) int {
	pPosition := started.Hour()*2 + 1
	if started.Minute() > 30 {
		pPosition = pPosition + 1
	}
	return pPosition
}
