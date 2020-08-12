package main

import (
	tool "github.com/glvd/go-media-tool"
	"github.com/glvd/ipcv/config"
	"github.com/glvd/ipcv/windows"
	"github.com/goextension/log"
	"github.com/goextension/log/zap"
)

func main() {
	zap.InitZapSugar()
	tool.RegisterLogger(log.Log())
	frame := windows.New(config.Mirror())
	frame.Run()
}
