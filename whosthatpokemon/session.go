package whosthatpokemon

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/google/uuid"
	"github.com/guregu/dynamo"
	"strings"
	"time"
)

type GameSession struct {
	SessionID      string    // The id for the session. Should be a randomly generated UUID
 	StartTime      time.Time // When the player started the game
 	CurrentPokemon *Pokemon  // The users Current Pokemon
 	Score int // The users current score for this session
	ExpirationTime time.Time // When this is removed from the session database
}

type StrippedGameSession struct {
	SessionID string
	CurrentPokemon *StrippedPokemon
	Score int
}

//NewGameSession Creates a new Game Session
func NewGameSession() (*GameSession, error) {
	id := uuid.New()
	newSession := &GameSession{
		SessionID:      id.String(),
		StartTime:      time.Now(),
		CurrentPokemon: newPokemon(),
		ExpirationTime: time.Now().Add(time.Hour * 6),  // Create a expiration time for this item.
	}
	err := newSession.save()
	if err != nil {
		return nil, err
	}
	return newSession, nil
}

func LoadGameSession(sessionID string) (*GameSession, error) {
	db := dynamo.New(session.New(), &aws.Config{Region:aws.String("us-east-1", )})
	table := db.Table(conf.SessionTable.TableName)

	var result *GameSession
	err := table.Get(conf.SessionTable.HashKey, sessionID).One(&result)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (gs *GameSession) NewStrippedSession() *StrippedGameSession{
	return &StrippedGameSession{
		SessionID:      gs.SessionID,
		CurrentPokemon: gs.CurrentPokemon.NewStrippedPokemon(),
		Score: gs.Score,
	}
}

func (gs *GameSession) save () error {
	db := dynamo.New(session.New(), &aws.Config{Region:aws.String("us-east-1", )})
	table := db.Table(conf.SessionTable.TableName)

	err := table.Put(gs).Run()

	if err != nil {
		return err
	}

	return nil
}

func (gs *GameSession) newPokemon() error {
	gs.CurrentPokemon = newPokemon() // Get a new pokemon
	err := gs.save()
	if err != nil {
		return err
	}
	return nil
}

func (gs *GameSession) CheckAnswer(answer string) (bool, error) {
	if strings.ToLower(answer) == gs.CurrentPokemon.Name {  // Check if their answer is the same as the current pokemon
		gs.incrementScore()
		err := gs.newPokemon()
		if err != nil {
			return false, err
		}
		return true, nil
	}
	return false, nil
}

// GameSession_incrementScore increments the score for a user
func (gs *GameSession) incrementScore() {
	gs.Score += 1
	return
}