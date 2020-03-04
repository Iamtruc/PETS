package main

import (
	"math/rand"
)

func read_circuit(my_circuit *TestCircuit){
	//Peerlist := my_circuit.Peers
	Inputlist := my_circuit.Inputs
	study_circuit := my_circuit.Circuit
	//Solution := my_circuit.ExpOutput
	var easy_input map[PartyID]uint64
	for i,j := range Inputlist{
		for _,secret := range j {
			easy_input[i] = secret
		}
	}
	read_gates(study_circuit, easy_input, 0)
}

func separateInShares(nbparty, a int) []int{

	var shares []int

	for i := 0;i<nbparty;i++{
		var e = rand.Intn(a+1)
		shares=append(shares, e)
		a=a-e
	}

	return shares
}