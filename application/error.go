package application

import (
	"fmt"
	"os"
)

func ExitWith(err error) {
	fmt.Fprintf(os.Stderr, "%s: %s\n", NAME, err)
}
