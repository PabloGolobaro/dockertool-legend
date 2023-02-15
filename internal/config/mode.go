package config

import "strconv"

type Mode struct {
	console bool
	port    int
}

func (m *Mode) Port() string {
	return strconv.Itoa(m.port)
}

func (m *Mode) Console() bool {
	return m.console
}

func NewMode(console bool, port int) *Mode {
	return &Mode{console: console, port: port}
}
