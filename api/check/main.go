package main

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/Beartime234/whos-that-pokemon/whosthatpokemon"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"log"
)

// Response is of type APIGatewayProxyResponse since we're leveraging the
// AWS Lambda Proxy Request functionality (default behavior)
//
// https://serverless.com/framework/docs/providers/aws/events/apigateway/#lambda-proxy-integration
type Response events.APIGatewayProxyResponse
type Request events.APIGatewayProxyRequest

type RequestBody struct {
	SessionID string  `json:"SessionID"`
	PokemonNameGuess string `json:"PokemonNameGuess"`
}

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(ctx context.Context, request Request) (Response, error) {
	var buf bytes.Buffer

	log.Printf("Request: %+v", request)

	var requestBody *RequestBody  // Unmarshal the body of the request
	err := json.Unmarshal([]byte(request.Body), &requestBody)

	session, err := whosthatpokemon.LoadGameSession(requestBody.SessionID)  // Create a new session

	if err != nil {
		return Response{StatusCode: 404}, err
	}

	_ = session.NewPokemon()  // TODO handle the error

	body, err := json.Marshal(session.NewStrippedSession())

	if err != nil {
		return Response{StatusCode: 404}, err
	}

	json.HTMLEscape(&buf, body)

	resp := Response{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            buf.String(),
		Headers: map[string]string{
			"Content-Type":			"application/json",
			"X-WTP-Func-Reply":		"api-Handler",
		},
	}
	return resp, nil
}

func main() {
	lambda.Start(Handler)
}
