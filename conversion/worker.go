package conversion

import (
	"context"
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
	Run()
	Status() string
	Stop()
	HookInfo(f func(s string))
}

type work struct {
	ctx      context.Context
	cancel   context.CancelFunc
	cfg      config.Conversion
	id       string
	status   *atomic.String
	filepath string
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

func (w *work) Run() {
	w.SetStatus(WorkStateRunning)
	ff := tool.NewFFMpeg()
	err := ff.Run(w.ctx, w.filepath)
	if err != nil {
		return
	}
}

func (w *work) Stop() {
	w.SetStatus(WorkStateStop)
}

func (w *work) HookInfo(f func(s string)) {
	if f != nil {
		f("write some string")
	}
}
