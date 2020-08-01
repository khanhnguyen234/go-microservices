package common

import (
	"github.com/fatih/color"
)

func LogErrorService(err error, msg string) {
	if err != nil {
		color.Red("FAIL: %s %s", msg, err)
	}
}

func LogSuccess(msg string) {
	color.Green("SUCCESS: %s", msg)
}