package main

import (
	"fmt"
	"sync"
	"testing"
)

func TestEval(t *testing.T){
	/// A test where we get three peers to communicate.
	for _,smart_circuit := range TestCircuits {
		peers := smart_circuit.Peers

		N := uint64(len(peers))
		P := make([]*LocalParty, N, N)
		dummyProtocol := make([]*DummyProtocol, N, N)
		// We create a new LocalParty and make the three peers communicate.
		var err error
		wg := new(sync.WaitGroup)
		for i := range peers {
			P[i], err = NewLocalParty(i, peers)
			P[i].WaitGroup = wg
			check(err)

			dummyProtocol[i] = P[i].NewDummyProtocol(uint64(i + 10))
		}

		network := GetTestingTCPNetwork(P)
		fmt.Println("parties connected")

		for i, Pi := range dummyProtocol {
			Pi.BindNetwork(network[i])
		}

		for _, p := range dummyProtocol {
			p.Add(1)
			go p.Run()
		}
		wg.Wait()

		for _, p := range dummyProtocol {
			fmt.Println(p, "completed with output", p.Output)
		}

		fmt.Println("test completed")
	}
}