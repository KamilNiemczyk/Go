package main

import "math/rand"

func remove(slice []SharkAttack, s int) []SharkAttack {
	return append(slice[:s], slice[s+1:]...)
}

func random10(attacks []SharkAttack) []SharkAttack {
	newtable := make([]SharkAttack, 10)
	for i := 0; i < 10; i++ {
		random := rand.Intn(len(attacks))
		newtable[i] = attacks[random]
		attacks = remove(attacks, random)
	}
	return newtable
}
