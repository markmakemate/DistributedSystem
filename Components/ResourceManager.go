package Worker

import "DistributedSystem/Zookeeper"

/*
ResourceMananger is the key component of this framework. It plays a role of system monitor, job scheduler,
resource scheduler.
When user execute the main() function to instantiate a job instance which is configured by user, system will give it

*/
type ResourceManager struct {
	uuid            uint64 //partition id
	timeStamp       uint64
	schedulerCenter *Zookeeper.SchedulerCenter
	jobClientMapper map[uint64]uint64         //<JId, CId>
	jobResultMapper map[uint64]AbstractResult //<JId, Result>
	jobQueue        chan *AbstractJob
	clientManager   *ClientManager
}

func (rm *ResourceManager) New(timeStamp *uint64, schedulerCenter *Zookeeper.SchedulerCenter,
	jobQueue *chan *AbstractJob) {
	&rm.timeStamp = timeStamp
	rm.schedulerCenter = schedulerCenter
	&rm.jobQueue = jobQueue
}

func (rm *ResourceManager) Rsync(path Path) {

}

func (rm *ResourceManager) Processor() {
	for job := range rm.jobQueue {

	}
}

func (rm *ResourceManager) GenerateUUID() *uint64 {

}
