package main

import (
	"github.com/google/uuid"
	"time"
)

type GameSession struct {
	ID             string    // The id for the session
 	StartTime      time.Time // When the player started the game
 	CurrentPokemon Pokemon   // Their Current Pokemon
	ExpirationTime time.Time // When this is removed from the session database
}

func NewGameSession() *GameSession {
	id := uuid.New()
	return &GameSession{
		ID:             id.String(),
		StartTime:      time.Time{},
		CurrentPokemon: Pokemon{},
		ExpirationTime: time.Time{},
	}
}