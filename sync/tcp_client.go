package sync

import (
	models "bitbucket.org/avanz/anotherPomodoro/model"
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
)

const SyncPort int = 1234

type IClient interface {
	GetRemotePomodoro() (models.CurrentPomodoro, error)
}

type TcpClient struct {
	net.Conn
}

func NewTcpClient(remoteAddress string) IClient {
	split := strings.Split(remoteAddress, ":")
	remotePort := SyncPort
	if len(split) > 1 {
		var err error
		remotePort, err = strconv.Atoi(split[1])
		if err != nil {
			log.Fatal(err)
		}
	}
	c, err := net.Dial("tcp", fmt.Sprintf("%s:%d", split[0], remotePort))
	if err != nil {
		log.Fatal(err)
	}
	return &TcpClient{c}
}
func (c *TcpClient) GetRemotePomodoro() (models.CurrentPomodoro, error) {
	fmt.Fprintf(c, "\n")
	message, err := bufio.NewReader(c).ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	var currentPomodoro = models.CurrentPomodoro{}
	err = json.Unmarshal([]byte(message), &currentPomodoro)
	return currentPomodoro, err
}
