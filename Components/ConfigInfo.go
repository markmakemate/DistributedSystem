package Worker

import "DistributedSystem/MapReduce"

type Config struct {
	mapper           MapReduce.Mapper
	reducer          MapReduce.Reducer
	outputKeyType    interface{}
	outputRecordType interface{}
}

func (c *Config) SetMapper(function MapReduce.Mapper) {
	c.mapper = function
}

func (c *Config) GetMapper() MapReduce.Mapper {
	return c.mapper
}

func (c *Config) SetReducer(function MapReduce.Reducer) {
	c.reducer = function
}
func (c *Config) SetOutputKeyType(Type interface{}) {
	c.outputKeyType = Type
}
func (c *Config) SetOutputRecordType(Type interface{}) {
	c.outputRecordType = Type
}
