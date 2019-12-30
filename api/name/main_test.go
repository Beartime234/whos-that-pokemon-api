package main

import (
	"encoding/json"
	"github.com/Beartime234/whos-that-pokemon/whosthatpokemon"
	"github.com/aws/aws-lambda-go/events"
	"log"
	"testing"
)

func TestHandler(t *testing.T) {
	const newUserName = "CharmanderIsTheBest"
	session, err := whosthatpokemon.NewGameSession()
	if err != nil {
		log.Fatal("Error in creating game session.")
	}
	requestBody, err := json.Marshal(NameRequestBody{
		SessionID:        session.SessionID,
		UserName:newUserName,
	})
	got, err := Handler(nil, Request{
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

	if err != nil {
		log.Fatal(err)
	}

	if got.StatusCode != 200 {
		log.Fatal("Status code was wrong. Error most likely occurred.")
	}

	// Confirm that the username has been changed so we need to reload the session
	session, err = whosthatpokemon.LoadGameSession(session.SessionID)  // Load the new session

	if err != nil {
		log.Print("Error loading the session")
		log.Fatal(err)
	}

	if session.UserName != newUserName {
		log.Fatalf("Username was not changed got %s wanted %s", session.UserName, newUserName)
	}
}