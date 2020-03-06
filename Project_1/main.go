package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)


// executes test circuit using operations detailed in operations.go
func main() {// We get the argument that we put in the command line
	prog := os.Args[0]
	fmt.Println(prog)
	// os.Args provides access to raw command-line arguments. Note that the first value in this slice is the path to the program, and os.Args[1:] holds the arguments to the program.
	args := os.Args[1:]
	fmt.Println(args)
	// If there are less then two arguments => panic
	if len(args) < 2 {
		fmt.Println("Usage:", prog, "[Party ID] [Input]")
		os.Exit(1)
	}

	// If there was an error in the party ID => panic
	partyID, errPartyID := strconv.ParseUint(args[0], 10, 64)
	if errPartyID != nil {
		fmt.Println("Party ID should be an unsigned integer")
		os.Exit(1)
	}

	// If there was an error in the party input => panic
	partyInput, errPartyInput := strconv.ParseUint(args[1], 10, 64)
	if errPartyInput != nil {
		fmt.Println("Party input should be an unsigned integer")
		os.Exit(1)
	}

	Client(PartyID(partyID), partyInput)
}


func Client(partyID PartyID, partyInput uint64) {

	//N := uint64(len(peers))
	peers := map[PartyID]string {
		0: "localhost:6660",
		1: "localhost:6661",
		2: "localhost:6662",
	}

	// Create a local party 
	lp, err := NewLocalParty(partyID, peers)
	check(err)

	// Create the network for the circuit
	network, err := NewTCPNetwork(lp)
	check(err)

	// Connect the circuit network 
	err = network.Connect(lp)
	check(err)
	fmt.Println(lp, "connected")
	<- time.After(time.Second) // Leave time for others to connect

	// Create a new circuit evaluation protocol 
	dummyProtocol := lp.NewDummyProtocol(partyInput)
	// Bind evaluation protocol to the network
	dummyProtocol.BindNetwork(network)

	// Evaluate the circuit
	dummyProtocol.Run()

	fmt.Println(lp, "completed with output", dummyProtocol.Output)
}
