package Worker

import (
	"DistributedSystem/RegisterCenter"
)

type AbstractScheduler interface {
	GetId() int64
	SetId(uuid int64)
	PushJob(job *[]byte) error
	Register(scheduler *AbstractScheduler)
	GenerateUUID() *int64
	GetJobById(jid *int64) (AbstractJob, error)
	New(uuid int64, RegisterCenter *RegisterCenter.SchedulerCenter, MAX *int64)
	//RegisterCenter:
	//1. A SchedulerCenter Object -> GlobalScheduler
	//2. nil -> LocalScheduler
}
