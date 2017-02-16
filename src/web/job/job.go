package job

import (
	"web/models"
	"time"
	"github.com/astaxie/beego"
	"fmt"
	"runtime/debug"
	"strconv"
	"web/service"
)

type Job struct {
	Id         int                                               	// 任务ID
	Name       string                                            	// 任务名称
	Task       *models.Job                                      	// 任务对象
	RunFunc    func(time.Duration) bool			// 执行函数，超时返回nil,true
	Status     int                                               	// 任务状态，大于0表示正在执行中
	Concurrent bool                                              	// 同一个任务是否允许并行执行
}

func (j *Job) Run() {
	if !j.Concurrent && j.Status > 0 {
		beego.Warn(fmt.Sprintf("任务[%d]上一次执行尚未结束，本次被忽略。", j.Id))
		return
	}

	defer func() {
		if err := recover(); err != nil {
			beego.Error(err, "\n", string(debug.Stack()))
		}
	}()

	if workPool != nil {
		workPool <- true
		defer func() {
			<-workPool
		}()
	}

	beego.Debug(fmt.Sprintf("开始执行任务: %d", j.Id))

	j.Status++
	defer func() {
		j.Status--
	}()

	timeout := time.Duration(time.Hour * 24)
	isTimeout:= j.RunFunc(timeout)

	if isTimeout {
		beego.Error("任务执行超时")
	}
}

func NewJobFromDb(j *models.Job,book *models.Book) (*Job,error) {
	job:=&Job{
		Id:j.Id,
		Name:j.Name,
		Task:j,
		Status:0,
		Concurrent:true,
	}

	job.RunFunc = func(timeout time.Duration) bool{
		done := make(chan bool,1)
		go func() {
			done <- true
		}()
		service.UpdateBook(book,"cache_book_"+strconv.Itoa(j.BookId))

		select {
		case <-time.After(timeout):
			beego.Warn(fmt.Sprintf("任务执行时间超过%d秒", int(timeout/time.Second)))
			go func() {
				<-done
			}()
			return false
		case ok := <-done:
			return ok
		}
	}
	return job,nil
}