package whosthatpokemon

import (
	"testing"
)

// Tests that we generate random pokemon
func TestGenerateRandomPokedexID(t *testing.T) {
	var genTimes = 100000
	for i := 1;  i<=genTimes; i++ {
		got := GenerateRandomPokedexID()
		if got < 1 && got > conf.MaxPokemon {
			t.Fail()
		}
	}
}

// Tests that we are getting everything correctly back. We are just checking that the columns correctly.
func TestGetRandomPokemon(t *testing.T) {
	// If any of these fails we probably are getting the wrong dynamo columns
	got := GetRandomPokemon()
	if got.PokedexID == 0 {
		t.Fail()
	}
	if got.OriginalImageUrl == "" {
		t.Fail()
	}
	if got.BWImageUrl == "" {
		t.Fail()
	}
	if got.Name == "" {
		t.Fail()
	}
}