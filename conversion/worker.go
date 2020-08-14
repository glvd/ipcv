package conversion

import (
	"context"
	"fmt"
	"github.com/glvd/go-media-tool"
	"github.com/glvd/ipcv/config"
	"github.com/google/uuid"
	"go.uber.org/atomic"
)

const (
	WorkStateRunning = "running"
	WorkStateStop    = "stop"
)

type Worker interface {
	ID() string
	Run(ctx context.Context)
	Status() string
	Stop()
	HookInfo(f func(s string))
}

type work struct {
	ctx         context.Context
	cancel      context.CancelFunc
	cfg         config.Conversion
	mediaConfig *tool.Config
	id          string
	status      *atomic.String
	filepath    string
	hook        func(msg string)
}

var _ Worker = &work{}

func NewWork(cfg config.Conversion, path string, id string) Worker {
	return &work{
		cfg:      cfg,
		id:       id,
		filepath: path,
		status:   atomic.NewString(WorkStateStop),
	}
}

func RandomWork(cfg config.Conversion, path string) Worker {
	id := uuid.Must(uuid.NewRandom()).String()
	return NewWork(cfg, path, id)
}

func (w work) ID() string {
	return w.id
}

func (w work) Status() string {
	return w.status.Load()
}

func (w *work) SetStatus(status string) {
	w.status.Store(status)
}

func (w *work) Run(ctx context.Context) {
	tool.DefaultMpegName = w.cfg.FFMPEG
	w.SetStatus(WorkStateRunning)
	w.ctx, w.cancel = context.WithCancel(ctx)
	defer w.cancel()
	if w.mediaConfig == nil {
		w.mediaConfig = tool.DefaultConfig()
	}
	ff := tool.NewFFMpeg(w.mediaConfig.ConfigOptions())
	ff.HandleMessage(w.messageCallback)
	err := ff.Run(w.ctx, w.filepath)
	if err != nil {
		return
	}
}

func (w *work) MediaConfig() *tool.Config {
	return w.mediaConfig
}

func (w *work) SetMediaConfig(mediaConfig *tool.Config) {
	w.mediaConfig = mediaConfig
}

func (w *work) Stop() {
	w.SetStatus(WorkStateStop)
	w.cancel()
}

func (w *work) messageCallback(message string) {
	fmt.Println(message)
	if w.hook != nil {
		w.hook(message)
	}
}

func (w *work) HookInfo(f func(s string)) {
	w.hook = f
}
