package main

import (
	"github.com/Beartime234/whos-that-pokemon/whosthatpokemon"
	"log"
	"testing"
)

func TestHandler(t *testing.T) {
	session, err := whosthatpokemon.NewGameSession()
	if err != nil {
		log.Fatal("Error in creating game session.")
	}
	got, _ := Handler(nil, &Request{
		SessionID:        session.SessionID,
		PokemonNameGuess: session.CurrentPokemon.Name,
	})
	if got.StatusCode != 200 {
		log.Fatal("Status code was wrong. Error most likely occurred.")
	}
}