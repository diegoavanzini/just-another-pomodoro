package sync

import (
	"bitbucket.org/avanz/anotherPomodoro/common"
	"bitbucket.org/avanz/anotherPomodoro/repository"
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"
)

type Listener struct {
	repository    repository.IPomodoroRepository
	sharedAddress string
	sharedPort    int
}

type IListener interface {
	Start()
}

func NewListener(repository repository.IPomodoroRepository) IListener {
	l := &Listener{repository: repository}
	timeShareAddressAndPort := ""
	err := repository.Read("settings", "timeShareAddressAndPort", &timeShareAddressAndPort)
	if err != nil {
		common.MainErrorListener <- err
	}
	split := strings.Split(timeShareAddressAndPort, ":")
	l.sharedAddress = split[0]
	if len(split) > 1 {
		l.sharedPort, err = strconv.Atoi(split[1])
		if err != nil {
			common.MainErrorListener <- err
		}
	}
	return l
}

func (l *Listener) Start() {
	countConnection := 0
	go func(listener *Listener) {
		if listener.sharedAddress == "" {
			listener.sharedAddress = "127.0.0.1"
		}
		if listener.sharedPort == 0 {
			listener.sharedPort = 1234
		}
		l, err := net.Listen("tcp4", fmt.Sprintf("%s:%d", listener.sharedAddress, listener.sharedPort))
		if err != nil {
			common.MainErrorListener <- err
			return
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
	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		//if err != nil {
		//	common.MainErrorListener <- err
		//}

		temp := strings.TrimSpace(netData)
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
