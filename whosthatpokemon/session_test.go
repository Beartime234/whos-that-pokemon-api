package whosthatpokemon

import (
	"log"
	"testing"
)

func TestGameSession_save(t *testing.T) {
	session := NewGameSession()
	err := session.save()

	if err != nil {
		log.Fatal(err)
	}
}

func TestNewGameSession(t *testing.T) {
	_ = NewGameSession()
}

func TestLoadGameSession(t *testing.T) {
	session := NewGameSession()
	err := session.save()
	if err != nil {
		log.Fatal("Error in saving or creating session, load game session was not the problem here")
	}
	got := LoadGameSession(session.SessionID)
	if got.SessionID != session.SessionID {
		t.Fail()
	}
}