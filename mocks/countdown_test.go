package mocks

import (
	"bytes"
	"testing"
)

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
	sleeper := &SpySleeper{}
	Countdown(buffer, sleeper)
	got := buffer.String()
	want := `3
2
1
Go!`
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
	if sleeper.Calls != 3 {
		t.Errorf("Wrong number of sleep calls")
	}
}
