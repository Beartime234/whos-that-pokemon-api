package whosthatpokemon

import (
	"log"
	"testing"
)

func TestGameSession_save(t *testing.T) {
	session, err := NewGameSession()
	if err != nil {
		log.Fatal("Failed creating new session not saving")
	}

	err = session.save()

	if err != nil {
		log.Fatal(err)
	}
}

func TestNewGameSession(t *testing.T) {
	_, err := NewGameSession()
	if err != nil {
		log.Fatal(err)
	}
}

func TestLoadGameSession(t *testing.T) {
	session, err := NewGameSession()
	if err != nil {
		log.Fatal("Failed creating new session not saving")
	}
	err = session.save()
	if err != nil {
		log.Fatal("Error in saving or creating session, load game session was not the problem here")
	}
	got := LoadGameSession(session.SessionID)
	if got.SessionID != session.SessionID {
		t.Fail()
	}
}