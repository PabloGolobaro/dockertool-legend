package config

import (
	"flag"
	"time"
)

type Flags struct {
	Stream   bool
	Duration time.Duration
}

func GetModeFlags() (Flags, error) {
	stream := flag.Bool("stream", false, "Get one-time stats or get stats stream")
	dur := flag.Int("dur", 0, "Duration of stream in seconds")

	flag.Parse()

	d := time.Duration(*dur)

	return Flags{Stream: *stream, Duration: time.Duration(time.Second * d)}, nil
}
