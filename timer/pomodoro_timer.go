package timer

import (
	"time"
)

type PomodoroTimer struct {
	maxValue                    time.Duration
	pauseListener, stopListener chan bool
	TimerValueListener          chan time.Duration
}

func NewPomodoroTimer(maxValue time.Duration, pauseListener, stopListener chan bool) PomodoroTimer {
	return PomodoroTimer{
		maxValue:      maxValue,
		pauseListener: pauseListener,
		stopListener:  stopListener,
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
			case <-pt.stopListener:
				return
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
