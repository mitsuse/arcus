package commands

import (
	"fmt"
	"os"
)

const (
	APP_NAME    = "bullet"
	APP_VERSION = "0.0.1"
	APP_DESC    = "A command-line tool to send text or files to Pushbullet."
)

func printError(err error) {
	fmt.Fprintf(os.Stderr, "%s: %s\n", APP_NAME, err)
}
