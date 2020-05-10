package widget

import (
	"testing"
	"time"

	_ "fyne.io/fyne/test"
	"github.com/stretchr/testify/assert"
)

func TestProgressBar_SetValue(t *testing.T) {
	pause := make(chan bool)
	alert := make(chan bool)
	bar := NewTimerProgressBar(10*time.Second, pause, alert)

	assert.Equal(t, time.Duration(0), bar.Min)
	assert.Equal(t, 10*time.Second, bar.Max)
	assert.Equal(t, time.Duration(0), bar.Value)

	bar.SetValue(5 * time.Second)
	assert.Equal(t, 5*time.Second, bar.Value)
}

func TestProgressBar_StartTimer(t *testing.T) {
	pause := make(chan bool)
	alert := make(chan bool)
	bar := NewTimerProgressBar(5*time.Second, pause, alert)

	assert.Equal(t, time.Duration(0), bar.Min)
	assert.Equal(t, 5*time.Second, bar.Max)
	assert.Equal(t, time.Duration(0), bar.Value)

	go bar.Start()
	a := <-alert

	assert.True(t, a)
	assert.True(t, bar.Value > time.Duration(0))

}
