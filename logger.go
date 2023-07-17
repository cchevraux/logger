package logger

import (
	"fmt"
)

var Version string = "1.0"

func log(mess string) {
	fmt.Println("[LOG] " + mess)
}
