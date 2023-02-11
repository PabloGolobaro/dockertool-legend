//go:build dev

package logger

import "os"

func init() {
	os.Setenv("MODE", "dev")
}
