package main

import (
	"fmt"
	"time"
)

type Process struct {
	ID       int
	Messages chan string
	IsLeader bool
}

func NewProcess(id int) *Process {
	return &Process{
		ID:       id,
		Messages: make(chan string),
		IsLeader: false,
	}
}

func (p *Process) bullyAlgorithm(processes []*Process) {
	// First assume me as the leader
	p.IsLeader = true

	for _, proc := range processes {
		// Sends "election" messages to all processes with higher IDs
		// If the process doesn't receive a "leader" message within a certain time window, it remains the leader
		if proc != nil && proc.ID > p.ID {
			proc.Messages <- "election"
			p.IsLeader = false
		}
	}

	time.Sleep(1 * time.Second)

	if !p.IsLeader {
		return
	}

	// Notify all processes with lower IDs that a new leader has been elected
	for _, proc := range processes {
		if proc != nil && proc.ID < p.ID {
			proc.Messages <- "leader"
		}
	}
}

func (p *Process) handleMessage(processes []*Process) {
	for {
		msg := <-p.Messages
		switch msg {
		case "election":
			go p.bullyAlgorithm(processes)
		case "leader":
			p.IsLeader = false
		}
	}
}

func main() {
	clusterSize := 5
	processes := make([]*Process, clusterSize)

	for i := 0; i < clusterSize; i++ {
		processes[i] = NewProcess(i + 1)
		go processes[i].handleMessage(processes)
	}

	processes[0].bullyAlgorithm(processes)

	time.Sleep(2 * time.Second)

	for _, proc := range processes {
		if proc.IsLeader {
			fmt.Printf("Process %d is the leader.\n", proc.ID)
		}
	}

	// Simulate an outage of process 5
	processes[4] = nil

	// Trigger a new election from process 3
	processes[2].bullyAlgorithm(processes)

	time.Sleep(2 * time.Second)

	for _, proc := range processes {
		if proc != nil && proc.IsLeader {
			fmt.Printf("Process %d is the new leader after outage.\n", proc.ID)
		}
	}
}
