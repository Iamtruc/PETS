package main

import(
	"math/rand"
)

// We have n is the number of parties that are communicating, and l is the modulus of the nb

func generatebeaver(nbparty,myrange int) (int, []int){
	var a []int
	var c int
	for i := 0;i<nbparty;i++{
		a = append(a, rand.Intn(myrange))
	}

	for _,i := range(a){
		c *= i
	}
	return c,a
}