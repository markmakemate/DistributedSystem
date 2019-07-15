package MapReduce

import Worker "DistributedSystem/Components"


/*
It's the implementations of mapreduce algorithm;

 */
type Mapper struct {
	key interface{}
	record *Worker.AbstractRecord
	method string
}


type Reducer struct {
	key interface{}
	recordList *[]Worker.AbstractRecord
	method string
}

