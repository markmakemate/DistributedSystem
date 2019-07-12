package Worker

type Job struct{
	JID int64
	cid int64
	Key string
	Result string
}
func (w *Job)New(id *int64, data *string)  {
	&w.Key = data
	&w.JID = id
}

func (w *Job) GetCId() int64 {
	return w.cid
}

func (w *Job) SetCId(cid int64)  {
	w.cid = cid
}
func (w *Job) GetJId() int64 {
	return w.JID
}
func (w *Job) SetJId(jid *int64)  {
	&w.JID = jid
}

