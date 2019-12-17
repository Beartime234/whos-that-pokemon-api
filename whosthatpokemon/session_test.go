package whosthatpokemon

import (
	"log"
	"testing"
)

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
	got, err := LoadGameSession(session.SessionID)
	if err != nil {
		panic(err)
	}
	if got.SessionID != session.SessionID {
		t.Fail()
	}
}

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

func TestGameSession_NewPokemon(t *testing.T) {
	session, err := NewGameSession()
	if err != nil {
		log.Fatal("Failed creating new session not saving")
	}
	firstPokemon :=  session.CurrentPokemon.Name
	err = session.NewPokemon()
	if err != nil {
		panic(err)
	}
	if firstPokemon == session.CurrentPokemon.Name {
		log.Printf("A new pokemon was not generated")
		t.Fail()
	}
}

func TestGameSession_Check(t *testing.T) {
	t.Skip() // TODO implement
}