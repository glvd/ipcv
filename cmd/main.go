package main

import (
	"github.com/glvd/ipcv/config"
	"github.com/glvd/ipcv/windows"
)

func main() {
	frame := windows.New(config.Load())
	frame.Run()
}
