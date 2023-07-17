package logger

import (
	"fmt"
)

var Version string = "1.1"

func Log(mess string) {
	fmt.Println("[LOG] " + mess)
}
