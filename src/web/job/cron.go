package job

import (
	"github.com/jakecoffman/cron"
	"sync"
	"github.com/astaxie/beego"
	"strconv"
)

var(
	mainCron *cron.Cron
	workPool chan bool
	lock sync.Mutex
)

func init() {
	if size,_:=beego.AppConfig.Int("jobs.pool"); size > 0{
		workPool=make(chan bool,size)
	}
	mainCron=cron.New()
	mainCron.Start()
}
///book/taskUpdate/6
func AddJob(spec string, job *Job) bool {
	lock.Lock()
	defer lock.Unlock()

	mainCron.AddJob(spec, job,strconv.Itoa(job.Id))
	return true
}

func RemoveJob(id string) {
	mainCron.RemoveJob(id)
}

func GetEntryById(id int) *cron.Entry {
	entries := mainCron.Entries()
	for _, e := range entries {
		if v, ok := e.Job.(*Job); ok {
			if v.Id == id {
				return e
			}
		}
	}
	return nil
}

func GetEntries(size int) []*cron.Entry {
	ret := mainCron.Entries()
	if len(ret) > size {
		return ret[:size]
	}
	return ret
}