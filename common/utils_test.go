package common

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestStringToDuration(t *testing.T) {
	duration, _ := StringToDuration("25:00")
	assert.Equal(t, 25*time.Minute, duration)

	duration, _ = StringToDuration("25:25")
	assert.Equal(t, 1525*time.Second, duration)
}
