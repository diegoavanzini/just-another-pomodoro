package common

import (
	"errors"
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"
)

func DurationToString(d time.Duration) string {
	d = d.Round(time.Second)
	h := d / time.Minute
	d -= h * time.Minute
	m := d / time.Second
	return fmt.Sprintf("%02d:%02d", h, m)
}

// TODO test
func StringToDuration(d string) (time.Duration, error) {
	var ret time.Duration
	minutes_seconds := strings.Split(d, ":")
	minutesFloat64, err := strconv.ParseFloat(minutes_seconds[0], 64)
	if err != nil {
		return ret, err
	}
	minutesDuration := time.Duration(minutesFloat64) * 60 * time.Second

	secondsFloat64, err := strconv.ParseFloat(minutes_seconds[1], 64)
	if err != nil {
		return ret, err
	}
	secondsDuration := time.Duration(secondsFloat64) * time.Second
	now := time.Now()
	ret = now.Add(secondsDuration).Add(minutesDuration).Sub(now)
	return ret, nil
}

func ExternalIP() (string, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}
		addrs, err := iface.Addrs()
		if err != nil {
			return "", err
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip == nil || ip.IsLoopback() {
				continue
			}
			ip = ip.To4()
			if ip == nil {
				continue // not an ipv4 address
			}
			return ip.String(), nil
		}
	}
	return "", errors.New("are you connected to the network?")
}
