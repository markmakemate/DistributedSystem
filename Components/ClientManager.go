package Worker

import (
	"DistributedSystem/Utils"
	"net/http"
)

type Client struct {
	clientInstance *http.Request
	cId            uint64
}

func (c *Client) New(clientInstance *http.Request) {
	c.clientInstance = clientInstance
}

type ClientManager struct {
	ClientMapper map[uint64]*Client
	ClientPool   chan *Client
}

func (cm *ClientManager) New(clientPool *chan *Client) {
	&cm.ClientPool = clientPool
}

func (cm *ClientManager) SetCId(client *Client) {
	//Client Id generation algorithm
	addr := Utils.IpParser(client.clientInstance.RemoteAddr)
	cid := new(uint64)
	for i := 0; i < len(addr); i++ {
		*cid += uint64(addr[i] << (len(addr) - i - 1))
	}
}

func (cm *ClientManager) FindClientById(cid *uint64) Client {
	return *cm.ClientMapper[*cid]
}
