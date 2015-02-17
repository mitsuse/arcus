/*
Package "app" provides the information of "bullet" command and some utilities.
*/
package app

import (
	"fmt"
	"os"
)

const (
	NAME         = "bullet"
	VERSION      = "0.1.0"
	DESC         = "A command-line tool to send text or files to Pushbullet."
	AUTHOR       = "Tomoya Kose (mitsuse)"
	AUTHOR_EMAIL = "tomoya@mitsuse.jp"
)

func PrintError(err error) {
	fmt.Fprintf(os.Stderr, "%s: %s\n", NAME, err)
}
