package Worker


/*

SchedulerCenter is a special component which store sessions of requests from different schedulers and can find a scheduler quickly
with a given uuid. On the other hand, SchedulerCenter also serves as the role of zookeeper to keep eventually consistency and load balance.
It implements a go version of Raft proxy and a distributed lock.

 */
type SchedulerCenter struct {
	SchedulerMapper map[uint64] *AbstractSchedulerSession
	SchedulerPool []*AbstractSchedulerSession
	results chan *AbstractResult
	job chan *AbstractJob
}

func (s *SchedulerCenter) Register(scheduler *AbstractSchedulerSession)  {
	s.SchedulerPool = append(s.SchedulerPool, scheduler)
	s.SchedulerMapper[(*scheduler).GetId()] = scheduler
}

func (s *SchedulerCenter)Deliver(job *AbstractJob)  {
	s.job <- job
}


