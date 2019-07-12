package Utils

import (
	"DistributedSystem/Components"
)

type Heap struct {
	workers []*Worker.AbstractScheduler
}
func (h *Heap) Parent(i int) int{

}
func (h *Heap) LeftChild(i int) int  {
	return 2*i-1
}
func (h *Heap) RightChild(i int) int  {
	return 2*i
}
func (h *Heap) Swap(i, j int) {
	h.workers[i], h.workers[j] = h.workers[j], h.workers[i]
}
func (h *Heap) Push(x *Worker.AbstractScheduler) {
	pos := len(h.workers)
	h.workers[pos] = x
	index := h.Parent(pos)
	for index != 0{
		if (*h.workers[index]).GetSchedulerInfo().GetLoader() > (*h.workers[pos]).GetSchedulerInfo().GetLoader(){
			h.Swap(index, pos)
			pos = index
			index = h.Parent(index)
		}else {
			break
		}
	}
}
func (h *Heap)Pop() {
	h.workers = h.workers[1:]
	h.Build(&h.workers)
}
func (h *Heap) Top() *Worker.AbstractScheduler{
	return h.workers[0]
}
func (h *Heap) Build(input *[]*Worker.AbstractScheduler)  {
	index := 0
	&h.workers = input
	for index != len(h.workers){
		left := h.LeftChild(index)
		right := h.RightChild(index)
		if (*h.workers[index]).GetSchedulerInfo().GetLoader() > (*h.workers[left]).GetSchedulerInfo().GetLoader(){
			h.Swap(index, left)
		}else if (*h.workers[index]).GetSchedulerInfo().GetLoader() > (*h.workers[right]).GetSchedulerInfo().GetLoader(){
			h.Swap(index, right)
		}
	}
}

func (h *Heap) Get() *[]*Worker.AbstractScheduler {
	return &h.workers
}
