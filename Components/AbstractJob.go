package Worker

type AbstractJob interface {
	New(id *int64, data *string)
	GetCId() int64
	SetCId(cid int64)
	GetJId() int64
	SetJId(jid *int64)
}
