package logger

import (
	"fmt"
)

var Version string = "1.2"

func Log(mess string) {
	fmt.Println("[LOG] " + mess)
}
