package Zookeeper

import "DistributedSystem/Components"

/*

SchedulerCenter is a special and critical component which store sessions of requests from different schedulers and can find a scheduler quickly
with a given uuid. On the other hand, SchedulerCenter also serves as the role of zookeeper to keep eventually consistency and load balance.
It implements a go version of Raft proxy and a distributed lock.

*/
type SchedulerCenter struct {
	SchedulerMapper map[uint64]*Worker.AbstractSchedulerSession
	SchedulerPool   []*Worker.AbstractSchedulerSession
	results         chan *Worker.AbstractResult
	jobChan         chan *Worker.AbstractJob
}

func (s *SchedulerCenter) RegisterScheduler(scheduler *Worker.AbstractSchedulerSession) {
	s.SchedulerPool = append(s.SchedulerPool, scheduler)
	s.SchedulerMapper[Worker.GetId()] = scheduler
}

func (s *SchedulerCenter) Map() {

}

func (s *SchedulerCenter) Register(job *Worker.AbstractJob) {
	s.jobChan <- job
}

func (s *SchedulerCenter) Reduce() {

}
