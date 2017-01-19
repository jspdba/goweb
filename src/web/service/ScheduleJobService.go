package service

import (
	"github.com/astaxie/beego"
	"runtime/debug"
	"sync/atomic"
	"github.com/jakecoffman/cron"
	"sync"
	"reflect"
)

const DEFAULT_JOB_POOL_SIZE = 10

var (
	// Singleton instance of the underlying job scheduler.
	MainCron *cron.Cron

	// This limits the number of jobs allowed to run concurrently.
	workPermits chan struct{}

	// Is a single job allowed to run concurrently with itself?
	selfConcurrent bool
)
func init() {
	MainCron = cron.New()
	workPermits = make(chan struct{}, DEFAULT_JOB_POOL_SIZE)
	selfConcurrent = false
	MainCron.Start()
}

type Job struct {
	Name    string
	inner   cron.Job
	status  uint32
	running sync.Mutex
}

const UNNAMED = "(unnamed)"

func New(job cron.Job) *Job {
	name := reflect.TypeOf(job).Name()
	if name == "Func" {
		name = UNNAMED
	}
	return &Job{
		Name:  name,
		inner: job,
	}
}

func (j *Job) Status() string {
	if atomic.LoadUint32(&j.status) > 0 {
		return "RUNNING"
	}
	return "IDLE"
}

func (j *Job) Run() {
	// If the job panics, just print a stack trace.
	// Don't let the whole process die.
	defer func() {
		if err := recover(); err != nil {
			beego.Error("%v", debug.Stack())
		}
	}()

	if !selfConcurrent {
		j.running.Lock()
		defer j.running.Unlock()
	}

	if workPermits != nil {
		workPermits <- struct{}{}
		defer func() { <-workPermits }()
	}

	atomic.StoreUint32(&j.status, 1)
	defer atomic.StoreUint32(&j.status, 0)

	j.inner.Run()
}