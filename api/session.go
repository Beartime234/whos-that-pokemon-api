package main

import "time"

type GameSession struct {
	Id string // The id for the session
 	StartTime time.Time // When the player started the game
 	CurrentPokemon Pokemon  // Their Current Pokemon
	ExpirationTime time.Time  // When this is removed from the session database
}