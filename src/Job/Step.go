package Job

import (
	"context"
	"time"
	"log"
)

type Step interface {
	Start()
	Stop()
	CanGoOn()bool
	SetCanGoOn(bool)
}

type Job struct {
	Id int
	Ctx context.Context
	Status int
	Cancel context.CancelFunc
	CanNext bool
}

func NewJob(ctx context.Context,cancelFunc context.CancelFunc)*Job{
	res:=&Job{}
	res.Id = getCount()
	res.Ctx = ctx
	res.Status = Initing
	res.Cancel = cancelFunc
	res.CanNext =false
	return res
}

const (
	Running int = iota
	Initing
	Complecated
	Failed
)
func (job *Job)Start(){

}
func(job *Job)Stop(){

}

func(job *Job)Wait(){
	for{
		time.Sleep(1*time.Second)
		select{
		case <-job.Ctx.Done():
			log.Printf("任务%d结束",job.Id)
			if job.Status!=Failed{
				log.Println("其他任务失败导致被停止了")
			}
			return
		default:
			if job.Status!=Running{
				job.Status = Running
			}
			log.Printf("Job %d is working.....",job.Id)
		}
	}
}

func(job *Job)GetStatus()int{
	return job.Status;
}
func(job *Job)CanGoOn()bool{
	return job.CanNext
}
func(job *Job)SetCanGoOn(goon bool){
	job.CanNext = goon
}

