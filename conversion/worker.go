package conversion

import (
	"context"
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
	ctx    context.Context
	cancel context.CancelFunc
	cfg    config.Conversion
	id     string
	status *atomic.String
}

var _ Worker = &work{}

func NewWork(cfg config.Conversion, id string) Worker {
	return &work{
		cfg:    cfg,
		id:     id,
		status: atomic.NewString(""),
	}
}

func RandomWork(cfg config.Conversion) Worker {
	id := uuid.Must(uuid.NewRandom()).String()
	return NewWork(cfg, id)
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
}

func (w *work) Stop() {
	w.SetStatus(WorkStateStop)
}

func (w *work) HookInfo(f func(s string)) {
	if f != nil {
		f("write some string")
	}
}
