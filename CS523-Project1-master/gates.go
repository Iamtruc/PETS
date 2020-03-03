package main

import (

)

func read_gates(my_circuit []Operation, param map[PartyID]uint64, my_ID PartyID){
	var dino map[WireID]uint64
	for _, opp := range my_circuit{
		switch opp.(type){
		case Input:
			dino[(opp.(Input)).Out] = param[(opp.(Input)).Party]
		case Add:
			dino[(opp.(Add)).Out] = dino[(opp.(Add)).In1] + dino[(opp.(Add)).In2]
		case Sub:
			dino[(opp.(Sub)).Out] = dino[(opp.(Sub)).In1] - dino[(opp.(Sub)).In2]
		case Mult:
			dino[(opp.(Mult)).Out] = dino[(opp.(Mult)).In1] * dino[(opp.(Mult)).In2]
		case MultCst:
			dino[(opp.(MultCst)).Out] = dino[(opp.(MultCst)).In] * (opp.(MultCst)).CstValue
		case AddCst:
			switch my_ID{
			case 0:// Only the first user car add the constant, the other just keep their value.
				dino[(opp.(AddCst)).Out] = dino[(opp.(AddCst)).In ] + (opp.(AddCst)).CstValue
			default:
				dino[(opp.(AddCst)).Out] = dino[(opp.(AddCst)).In ]
			}
		case Reveal:
			// Broadcast result
		}
	}
}