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
	works     *sync.Pool
	workStore *sync.Map
	running   *atomic.Bool
	routines  *atomic.Int32
	MaxLimit  int32
}

// DefaultLimit ...
var DefaultMaxLimit int32 = 2
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
		if t.routines.Load() < t.MaxLimit {
			fmt.Println("running:", t.routines.Load(), "max:", t.MaxLimit)
			work := t.getWork()
			if work != nil {
				t.routines.Add(1)
				go t.startWork(t.ctx, work)
			} else {
				fmt.Println("task running")
				time.Sleep(3 * time.Second)
			}
		} else {
			fmt.Println("task running,no limit sleeping")
			time.Sleep(5 * time.Second)
			continue
		}
	}
}

func Start() error {
	return _task.Start()
}

func (t *Task) IsRunning() bool {
	return t.running.Load()
}

// Start ...
func (t *Task) Start() error {
	if t.running.CAS(false, true) {
		t.ctx, t.cancel = context.WithCancel(context.TODO())
		go t.run()
	}
	return nil
}

// StopWork ...
func (t *Task) StopWork(id string) {
	worker := t.GetWorker(id)
	if worker.Status() == WorkStateRunning {
		worker.Stop()
	}
}

func (t *Task) checkWork(work Worker) error {
	//check work rule
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
		fmt.Println("new work", v.(Worker).ID())
		return v.(Worker)
	}
	return nil
}

func (t *Task) startWork(ctx context.Context, worker Worker) {
	defer t.routines.Add(-1)
	time.Sleep(5 * time.Second)
	worker.Run()
}

// NewTask ...
func NewTask() *Task {
	return &Task{
		works:     &sync.Pool{},
		workStore: &sync.Map{},
		running:   atomic.NewBool(false),
		routines:  atomic.NewInt32(0),
		MaxLimit:  DefaultMaxLimit,
	}
}
