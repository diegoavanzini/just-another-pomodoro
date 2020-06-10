package timer

import (
	"time"
)

type PomodoroTimer struct {
	maxValue           time.Duration
	pauseListener      chan bool
	TimerValueListener chan time.Duration
}

func NewPomodoroTimer(maxValue time.Duration, pauseListener chan bool) PomodoroTimer {
	return PomodoroTimer{
		maxValue:      maxValue,
		pauseListener: pauseListener,
	}
}

func (pt *PomodoroTimer) StartTimer(doSomething func(value time.Duration)) {
	ticker := time.NewTicker(1 * time.Second)
	value := pt.maxValue
	func() {
		for {
			select {
			case p := <-pt.pauseListener:
				if p {
					<-pt.pauseListener
				}
			case <-ticker.C:
 				value -= 1 * time.Second
				if value < 0 {
					return
				}
				doSomething(value)
			}
		}
	}()
}
