package Worker

//Abstraction of a Job
//Users can implement a Job Object by themselves, or they can call Job struct directly
type AbstractJob interface {
	New(id *uint64, ctx AbstractTaskContext) interface{}

	GetCId() uint64
	SetCId(cid *uint64)
	GetJId() uint64
	SetJId(jid *uint64)

	SetTaskContext(ctx AbstractTaskContext)
	GetTaskContext() AbstractTaskContext
	Submit()
	Do() AbstractResult
}

//Abstract interface of a Scheduler which runs locally
//Users can implement it by themselves
type AbstractScheduler interface {
	GetId() uint64
	SetId(uuid *uint64)
	PushJob(job *[]byte) error
	GetJobById(jid *uint64) (AbstractJob, error)
	New(uuid *uint64, jobQueue *chan *AbstractJob)
}

//Store the session of a request from another machine
type AbstractSchedulerSession interface {
	GetId() uint64
	SetId(uuid *uint64)
	Close()
	Ping() bool
	AssignJob(job AbstractJob) error
	GetResult(jid *uint64) (AbstractResult, error)
	GetStatus() Status
}

//All config information of one Job
type AbstractTaskContext interface {
	GetJId() uint64
	SetKey(key interface{})
	SetRecord(record interface{})
	GetConfig() Config
	SetConfig(conf Config)
}

type AbstractRecord interface {
}

type AbstractResult interface {
}
