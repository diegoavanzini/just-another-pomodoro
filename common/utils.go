package common

import (
	"fmt"
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
