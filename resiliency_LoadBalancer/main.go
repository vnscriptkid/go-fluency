package main

import (
	"fmt"

	"github.com/vnscriptkid/go-fluency/resiliency_LoadBalancer/lib"
)

func main() {
	servers := []*lib.Server{
		{Host: "Server 1", Weight: 1},
		{Host: "Server 2", Weight: 2},
		{Host: "Server 3", Weight: 3},
	}

	loadBalancer := lib.NewLoadBalancer(lib.NewRoundRobin(servers))

	// Test Round Robin
	fmt.Println("Round Robin:")
	for i := 0; i < 6; i++ {
		server := loadBalancer.GetServer()
		fmt.Printf("Request %d -> %s\n", i+1, server.Host)
	}
}
