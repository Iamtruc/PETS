package main

import (

)

func read_circuit(my_circuit []*TestCircuit){
	for _, test := range my_circuit{
		my_peers := test.Peers
		// Do Beaver stuff, at the end of the Beaver stuff, you need to have param map[PartyID]uint64, my_ID GateID
		my_param := test.Inputs
		// We need new parameters
		circuit := test.Circuit
		for i :=0 ; i< len(my_peers); i++{
			var best_param map[PartyID]uint64
			var my_ID PartyID // How do we find the right PartyID ?
			read_gates(circuit, best_param,my_ID)
		}
	}


}
