package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/google/uuid"
	"github.com/guregu/dynamo"
	"time"
)

type GameSession struct {
	SessionID      string    // The id for the session. Should be a randomly generated UUID
 	StartTime      time.Time // When the player started the game
 	CurrentPokemon *Pokemon  // Their Current Pokemon
	ExpirationTime time.Time // When this is removed from the session database
}

//NewGameSession Creates a new Game Session
func NewGameSession() *GameSession {
	id := uuid.New()
	return &GameSession{
		SessionID:      id.String(),
		StartTime:      time.Now(),
		CurrentPokemon: NewPokemon(),
		ExpirationTime: time.Now().Add(time.Hour * 6),  // Create a expiration time for this in 6 hours
	}
}

func (gs *GameSession) save () error {
	db := dynamo.New(session.New(), &aws.Config{Region:aws.String("us-east-1", )})
	table := db.Table(SessionTableName)

	err := table.Put(gs).Run()

	if err != nil {
		return err
	}

	return nil
}