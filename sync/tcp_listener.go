package sync

import (
	"bitbucket.org/avanz/anotherPomodoro/common"
	"bitbucket.org/avanz/anotherPomodoro/repository"
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"strings"
	"time"
)

type Listener struct {
	repository repository.IPomodoroRepository
}

type IListener interface {
	Start()
}

func NewListener(repository repository.IPomodoroRepository) IListener {
	return &Listener{repository: repository}
}

func (l *Listener)Start() {
	countConnection := 0
	go func(listener *Listener) {
		l, err := net.Listen("tcp4", fmt.Sprintf("%s:%d", "127.0.0.1", SyncPort))
		if err != nil {
			common.MainErrorListener <- err
		}
		defer l.Close()

		for {
			c, err := l.Accept()
			if err != nil {
				common.MainErrorListener <- err
				return
			}
			go listener.handleConnection(c)
			countConnection++
		}
	}(l)
}

func (l *Listener) handleConnection(c net.Conn) {
	fmt.Print(".")
	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			common.MainErrorListener <- err
		}

		temp := strings.TrimSpace(string(netData))
		if temp == "STOP" {
			break
		}
		var currentTimerValue string
		err = l.repository.Read("current", "timerValue", &currentTimerValue)
		if err != nil {
			common.MainErrorListener <- err
		}
		var timeDuration = int((25 * time.Minute).Seconds())
		err = l.repository.Read("settings", "timeDuration", &timeDuration)
		if err != nil {
			common.MainErrorListener <- err
		}
		var pauseDuration = int((5 * time.Minute).Seconds())
		err = l.repository.Read("settings", "pauseDuration", &pauseDuration)
		if err != nil {
			common.MainErrorListener <- err
		}
		currentPomodoro := struct {
			TimeDuration      int
			PauseDuration     int
			CurrentTimerValue string
		}{timeDuration, pauseDuration, currentTimerValue}
		currentPomodoroJson, err := json.Marshal(currentPomodoro)
		if err != nil {
			common.MainErrorListener <- err
		}
		c.Write([]byte(string(currentPomodoroJson) + "\n"))
	}
	c.Close()
}
