package main

import (
	"testing"
)

// Tests that we generate random pokemon
func TestGenerateRandomPokedexID(t *testing.T) {
	var genTimes = 100000
	for i := 1;  i<=genTimes; i++ {
		got := GenerateRandomPokedexID()
		if got < 1 && got > MaxPokemon {
			t.Fail()
		}
	}
}

func TestGetRandomPokemon(t *testing.T) {
	got := GetRandomPokemon()
	if got.PokedexID == 0 {
		t.Fail()  // If we failed here we probably getting wrong column
	}
}