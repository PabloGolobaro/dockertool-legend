package app

import "time"

type Mode struct {
	stream    bool
	duration  time.Duration
	WithTimer bool
}

func NewMode(stream bool, duration time.Duration) Mode {
	withTimer := false
	if duration > 0 {
		withTimer = true
	}
	return Mode{stream: stream, duration: duration, WithTimer: withTimer}
}
