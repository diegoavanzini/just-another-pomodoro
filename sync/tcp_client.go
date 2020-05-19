package sync

import (
	"bitbucket.org/avanz/anotherPomodoro/common"
	models "bitbucket.org/avanz/anotherPomodoro/model"
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"strconv"
	"strings"
)

const SyncPort int = 1234

type IClient interface {
	GetRemotePomodoro() (models.CurrentPomodoro, error)
}

type TcpClient struct {
	c net.Conn
}

func NewTcpClient(remoteAddress string) IClient {
	split := strings.Split(remoteAddress, ":")
	remotePort := SyncPort
	if len(split) > 1 {
		var err error
		remotePort, err = strconv.Atoi(split[1])
		if err != nil {
			common.MainErrorListener <- err
		}
	}
	c, err := net.Dial("tcp", fmt.Sprintf("%s:%d", split[0], remotePort))
	if err != nil {
		common.MainErrorListener <- err
		return nil
	}
	return &TcpClient{c}
}
func (tc *TcpClient) GetRemotePomodoro() (models.CurrentPomodoro, error) {
	fmt.Fprintf(tc.c, "\n")
	message, err := bufio.NewReader(tc.c).ReadString('\n')
	if err != nil {
		common.MainErrorListener <- err
	}
	var currentPomodoro = models.CurrentPomodoro{}
	err = json.Unmarshal([]byte(message), &currentPomodoro)
	return currentPomodoro, err
}
