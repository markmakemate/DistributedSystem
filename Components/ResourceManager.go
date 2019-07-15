package Worker


/*
ResourceMananger is the key component of this framework. It plays a role of system monitor, job scheduler,
resource scheduler.
When user execute the main() function to instantiate a job instance which is configured by user, system will give it

 */
type ResourceManager struct {
	timeStamp uint64
	schedulerCenter *SchedulerCenter
	JobClientMapper map[uint64] uint64
	JobResultMapper map[uint64] interface{}
	JobQueue chan *AbstractJob
	clientManager *ClientManager
}

func (rm *ResourceManager) New(timeStamp *uint64, schedulerCenter *SchedulerCenter,
	jobQueue *chan *AbstractJob) {
	&rm.timeStamp = timeStamp
	rm.schedulerCenter = schedulerCenter
	&rm.JobQueue = jobQueue
}

func (rm *ResourceManager) Rsync(path Path)  {
	
}

func (rm *ResourceManager) Processor()  {
	for job := range rm.JobQueue{
		rm.schedulerCenter.Deliver(job)

	}
}

func (rm *ResourceManager)GenerateUUID() *uint64 {

}
