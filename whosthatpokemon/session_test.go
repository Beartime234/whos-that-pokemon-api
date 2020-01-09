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

func TestGameSession_newPokemon(t *testing.T) {
	session, err := NewGameSession()
	if err != nil {
		log.Fatal("Failed creating new session not saving")
	}
	firstPokemon :=  session.CurrentPokemon.Name
	err = session.newPokemon()
	if err != nil {
		panic(err)
	}
	if firstPokemon == session.CurrentPokemon.Name {
		log.Printf("A new pokemon was not set")
		t.Fail()
	}
}

func TestGameSession_Check(t *testing.T) {
	session, err := NewGameSession()
	if err != nil {
		log.Fatal("Failed creating new session not saving")
	}
	firstPokemon :=  session.CurrentPokemon.Name

	wasCorrect, err := session.CheckAnswer("bleh")
	if wasCorrect == true {
		log.Printf("Was wrong but returned true")
		t.Fail()
	}
	if firstPokemon != session.CurrentPokemon.Name {
		log.Printf("The answer was wrong and a new pokemon was set")
		t.Fail()
	}
	wasCorrect, err = session.CheckAnswer(firstPokemon)
	if wasCorrect == false {
		log.Printf("Was correct but returned false")
		t.Fail()
	}
	if firstPokemon == session.CurrentPokemon.Name {
		log.Printf("The answer was correct and a new pokemon wasn't set")
		t.Fail()
	}
}

func TestGameSession_CheckAnswer(t *testing.T) {
	session, err := NewGameSession()
	if err != nil {
		log.Fatal("Failed creating new session not CheckAnswer")
	}
	correct, err := session.CheckAnswer(session.CurrentPokemon.Name)
	if correct != true {
		t.Fail()
	}

	session.CurrentPokemon.Name = "bulbasaur"
	correct, err = session.CheckAnswer("bulbasaurr")
	if correct != true {
		t.Fail()
	}
}

func TestGameSession_isGuessCorrect(t *testing.T) {
	session, err := NewGameSession()
	if err != nil {
		log.Fatal("Failed creating new session not isGuessCorrect")
	}

	// Check that if your correct your correct
	session.CurrentPokemon.Name = "bulbasaur"
	correct, err := session.CheckAnswer(session.CurrentPokemon.Name)
	if correct == false {
		log.Printf("The answer was correct but was said to be wrong")
		t.Fail()
	}

	// Check that if your wrong your wrong
	session.CurrentPokemon.Name = "bulbasaur"
	correct, err = session.CheckAnswer("dog")
	if correct == true {
		log.Printf("The answer was wrong but was said to be correct")
		t.Fail()
	}

	// Check that if your close
	session.CurrentPokemon.Name = "bulbasaur"
	correct, err = session.CheckAnswer("bulbasaurr")
	if correct == false {
		log.Printf("The answer was close enough but was marked incorrect")
		t.Fail()
	}
}

func TestGameSession_SetUserName(t *testing.T) {
	session, err := NewGameSession()

	if err != nil {
		log.Fatal("Failed creating new session not SetUserName")
	}

	err = session.SetUserName("CharmanderLover")

	if err != nil {
		log.Fatal("Failed setting new username")
	}

	if session.UserName != "CharmanderLover" {
		log.Fatalf("Session username was %s wanted CharmanderLover", session.UserName)
	}
}

func Test_generateDefaultUserName(t *testing.T) {
	got := generateDefaultUserName()
	if got != "User-12345" {
		log.Printf("Got %s wanted User-12345", got)
		t.Fail()
	}
}