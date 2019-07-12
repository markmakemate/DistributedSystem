package Worker

import (
	"DistributedSystem/RegisterCenter"
	"encoding/json"
	"net"
)

type Scheduler struct {
	uuid int64
	SchedulerCenter *RegisterCenter.SchedulerCenter
	JobPool map[int64] *[]byte
	JobChan chan *[]byte
	ResultChan chan *[]byte
}
func (g *Scheduler) New(uuid int64,
	SchedulerCenter *RegisterCenter.SchedulerCenter, MAX *int64){
	g.uuid = uuid
	g.SchedulerCenter = SchedulerCenter
	g.JobChan = make(chan *[]byte, *MAX)
}

func (g *Scheduler)SetId(uuid int64)  {
	g.uuid = uuid
}

func (g *Scheduler)GetId() int64 {
	return g.uuid
}

func (g *Scheduler)GenerateUUID() *int64{
	if g.SchedulerCenter != nil{
		uuid := new(int64)
		return uuid
	}else {
		return nil
	}
}

func (g *Scheduler)Register(scheduler *AbstractScheduler){
	g.SchedulerCenter.Register(scheduler)
}

func (g *Scheduler) PushJob(job *[]byte) error {
	g.JobChan <- job
	var s Job
	err := json.Unmarshal(*job, &s)
	g.JobPool[s.JID] = job
	return err
}

func (g *Scheduler) GetJobById(jid *int64) (AbstractJob, error) {
	var s Job
	err := json.Unmarshal(*g.JobPool[*jid], &s)
	return &s, err
}

func (g *Scheduler) Send(conn net.Conn) error {
	job, ok :=<- g.JobChan
	if !ok{
		_,err := conn.Write(*job)
		return err
	}else{
		return nil
	}

}

