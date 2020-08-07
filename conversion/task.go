package conversion

import (
	"context"
	"fmt"
	"go.uber.org/atomic"
	"sync"
	"time"
)

// Task ...
type Task struct {
	ctx       context.Context
	cancel    context.CancelFunc
	works     sync.Pool
	workStore sync.Map
	running   *atomic.Bool
	routines  *atomic.Int32
	MaxLimit  int32
}

// DefaultLimit ...
var DefaultMaxLimit int32 = 3
var _task *Task

func init() {
	_task = NewTask()
}

func AddWorker(worker Worker) {
	_task.AddWorker(worker)
}

// AddWorker ...
func (t *Task) AddWorker(worker Worker) {
	if err := t.checkWork(worker); err != nil {
		return
	}
	t.workStore.Store(worker.ID(), worker)
	t.works.Put(worker)
}

func Stop() {
	_task.Stop()
}

// Stop ...
func (t *Task) Stop() {
	if t.running.CAS(true, false) {
		if t.cancel != nil {
			t.cancel()
		}
	}
}

func (t *Task) run() {
	for {
		if !t.running.Load() {
			return
		}
		if t.routines.Load() <= t.MaxLimit {
			work := t.getWork()
			if work != nil {
				t.routines.Add(1)
				go t.startWork(work)
			} else {
				fmt.Println("task running")
				time.Sleep(3 * time.Second)
			}
		} else {
			fmt.Println("task running,no work sleeping")
			time.Sleep(30 * time.Second)
			continue
		}
	}
}

func Start() error {
	return _task.Start()
}

// Start ...
func (t *Task) Start() error {
	if t.running.CAS(false, true) {
		t.ctx, t.cancel = context.WithCancel(context.TODO())
	}
	go t.run()
	return nil
}

// StopWork ...
func (t *Task) StopWork(id string) {

}

func (t *Task) checkWork(work Worker) error {
	return nil
}

func GetWorker(id string) Worker {
	return _task.GetWorker(id)
}

func (t *Task) GetWorker(id string) Worker {
	load, ok := t.workStore.Load(id)
	if ok {
		return load.(Worker)
	}
	return nil
}

func (t *Task) getWork() Worker {
	if v := t.works.Get(); v != nil {
		return v.(Worker)
	}
	return nil
}

func (t *Task) startWork(worker Worker) {

	defer t.routines.Add(-1)
	worker.Run()
}

// NewTask ...
func NewTask() *Task {
	return &Task{
		works:     sync.Pool{},
		workStore: sync.Map{},
		running:   atomic.NewBool(false),
		routines:  atomic.NewInt32(0),
		MaxLimit:  DefaultMaxLimit,
	}
}
