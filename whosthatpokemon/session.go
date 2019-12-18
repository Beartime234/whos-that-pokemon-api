package whosthatpokemon

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/google/uuid"
	"github.com/guregu/dynamo"
	"strings"
	"time"
)

// GameSession this is the object that controls the flow of the game
type GameSession struct {
	SessionID      string    // The id for the session. Should be a randomly generated UUID
 	StartTime      time.Time // When the player started the game
 	CurrentPokemon *Pokemon  // The users Current Pokemon
 	Score int // The users current score for this session
	ExpirationTime time.Time // When this is removed from the session database
}

// StrippedGameSession this is what is sent back to the application. It is stripped so users cannot see the
// name and other things that give away the answer.
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
		Score:0,
	}
	err := newSession.save()
	if err != nil {
		return nil, err
	}
	return newSession, nil
}

// LoadGameSession Loads a session from the database. Does not create a new session and must have a session id
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

// GameSession_NewStrippedSession creates a stripped session that you can return to the user
func (gs *GameSession) NewStrippedSession() *StrippedGameSession{
	return &StrippedGameSession{
		SessionID:      gs.SessionID,
		CurrentPokemon: gs.CurrentPokemon.NewStrippedPokemon(),
		Score: gs.Score,
	}
}

// This function checks if the answer
func (gs *GameSession) CheckAnswer(answer string) (bool, error) {
	if strings.ToLower(answer) == gs.CurrentPokemon.Name {  // Check if their answer is the same as the current pokemon
		// They were correct
		gs.incrementScore() // Increment the score
		err := gs.newPokemon()  // Generate a new pokemon
		if err != nil {
			return false, err
		}
		err = gs.save() // Save the new pokemon
		if err != nil {
			return false, err
		}
		return true, nil
	}
	return false, nil
}


// GameSession_save saves the game session to the database. This should be called whenever it's updated.
func (gs *GameSession) save () error {
	db := dynamo.New(session.New(), &aws.Config{Region:aws.String("us-east-1", )})
	table := db.Table(conf.SessionTable.TableName)

	err := table.Put(gs).Run()

	if err != nil {
		return err
	}

	return nil
}

// GameSession_newPokemon Generates a new pokemon for the current session
// NOTE: You would still need to save this
func (gs *GameSession) newPokemon() error {
	gs.CurrentPokemon = newPokemon() // Get a new pokemon
	return nil
}


// GameSession_incrementScore increments the score for the session
// NOTE: You would need to still save this
func (gs *GameSession) incrementScore() {
	gs.Score += 1
	return
}