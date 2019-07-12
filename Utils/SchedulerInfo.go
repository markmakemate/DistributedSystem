package Utils

type SchedulerInfo struct {
	loader float32
	uuid string
}

func (s *SchedulerInfo)GetLoader() float32 {
	return s.loader
}
