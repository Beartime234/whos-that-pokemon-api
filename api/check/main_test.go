package main

import (
	"encoding/json"
	"github.com/Beartime234/whos-that-pokemon/whosthatpokemon"
	"github.com/aws/aws-lambda-go/events"
	"log"
	"testing"
)

func TestHandler(t *testing.T) {
	session, err := whosthatpokemon.NewGameSession()
	if err != nil {
		log.Fatal("Error in creating game session.")
	}
	requestBody, err := json.Marshal(CheckRequestBody{
		SessionID:        session.SessionID,
		PokemonNameGuess: session.CurrentPokemon.Name,
	})
	got, _ := Handler(nil, Request{
		Resource:                        "",
		Path:                            "",
		HTTPMethod:                      "",
		Headers:                         nil,
		MultiValueHeaders:               nil,
		QueryStringParameters:           nil,
		MultiValueQueryStringParameters: nil,
		PathParameters:                  nil,
		StageVariables:                  nil,
		RequestContext:                  events.APIGatewayProxyRequestContext{},
		Body:                            string(requestBody),
		IsBase64Encoded:                 false,
	})
	if got.StatusCode != 200 {
		log.Fatal("Status code was wrong. Error most likely occurred.")
	}

	reloadedSession, err := whosthatpokemon.LoadGameSession(session.SessionID)

	if err != nil {
		log.Print("Error loading the session")
		log.Fatal(err)
	}

	if reloadedSession.CurrentPokemon == session.CurrentPokemon {
		t.Fatal("The new pokemon is the same as the old pokemon. Unless you are really unlucky there is no way" +
			"this would fail unless the code is wrong")
	}
}