package main

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