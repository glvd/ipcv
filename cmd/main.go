package main

import (
	"github.com/glvd/ipcv/config"
	"github.com/glvd/ipcv/windows"
)

func main() {
	env()
	frame := windows.New(config.Mirror())
	frame.Run()
}
