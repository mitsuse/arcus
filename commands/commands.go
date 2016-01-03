/*
Package "commands" provides the information of `arcus` command and sub-commands.
*/
package commands

import (
	"fmt"
	"os"
)

const (
	NAME         = `arcus`
	VERSION      = "0.1.3"
	DESC         = "A command-line tool to send a message to devices via Pushbullet."
	AUTHOR       = "Tomoya Kose (mitsuse)"
	AUTHOR_EMAIL = "tomoya@mitsuse.jp"
)

func printError(err error) {
	fmt.Fprintf(os.Stderr, "%s: %s\n", NAME, err)
}
