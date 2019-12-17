package main

import (
	"github.com/google/uuid"
	"time"
)

type GameSession struct {
	ID             string    // The id for the session. Should be a randomly generated UUID
 	StartTime      time.Time // When the player started the game
 	CurrentPokemon *Pokemon   // Their Current Pokemon
	ExpirationTime time.Time // When this is removed from the session database
}

//NewGameSession Creates a new Game Session
func NewGameSession() *GameSession {
	id := uuid.New()
	return &GameSession{
		ID:             id.String(),
		StartTime:      time.Now(),
		CurrentPokemon: NewPokemon(),
		ExpirationTime: time.Now().Add(time.Hour * 6),  // Create a expiration time for this in 6 hours
	}
}