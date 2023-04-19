package lib

import "log"

func main() {
	log.Println("Starting the application...")
}

type Client struct {
	Host string
	IP   string
}

type Balancer interface {
	Next(client *Client) *Server
}

type LoadBalancer struct {
	balancer Balancer
}

func NewLoadBalancer(balancer Balancer) *LoadBalancer {
	return &LoadBalancer{
		balancer: balancer,
	}
}

func (l *LoadBalancer) SetBalancer(balancer Balancer) {
	l.balancer = balancer
}

func (l *LoadBalancer) GetServer() *Server {
	return l.balancer.Next(nil)
}

type Server struct {
	Host   string
	Weight int
}
