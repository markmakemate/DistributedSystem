package Utils

import Worker "DistributedSystem/Components"

/*
Go version of SnowFlake - an uuid generating algorithm
*/
func SnowFlake(rm *Worker.ResourceManager) *uint64 {
	uuid := new(uint64)
	return uuid
}
func ClientIdGenerator(addr string) uint64 {
	result := new(uint64)

}
func IpParser(ip string) []uint8 {
	var result []uint8
	var candidate uint8
	candidate = 0
	for i := 0; i < len(ip); i++ {
		switch ip[i] {
		case '.':
			result = append(result, candidate)
		case ':':
			break
		default:
			candidate = candidate*10 + uint8(ip[i])
		}
	}
	return result
}
