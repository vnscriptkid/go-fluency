package lib

type RoundRobin struct {
	servers []*Server
	index   int
}

func NewRoundRobin(servers []*Server) *RoundRobin {
	return &RoundRobin{
		servers: servers,
		index:   0,
	}
}

func (r *RoundRobin) Next(client *Client) *Server {
	server := r.servers[r.index]
	r.index = (r.index + 1) % len(r.servers)
	return server
}
