package RegisterCenter

import (
	"DistributedSystem/Utils"
	"DistributedSystem/Components"
)

type SchedulerCenter struct {
	SchedulerPool Utils.Heap
	uuid *[]int64
}
//Register为原子操作
func (s *SchedulerCenter) Register(scheduler *Worker.AbstractScheduler)  {
	s.SchedulerPool.Push(scheduler)
}
//Get为原子操作
func (s *SchedulerCenter) Get() *Worker.AbstractScheduler {
	
}
