package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/Beartime234/whos-that-pokemon/whosthatpokemon"
	goaway "github.com/TwinProduction/go-away"
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

type NameRequestBody struct {
	SessionID string  `json:"SessionID"`
	UserName string `json:"UserName"`
}

type NameResponseBody struct {
	Session *whosthatpokemon.StrippedGameSession
}

func NewStartResponseBody(session *whosthatpokemon.GameSession) *NameResponseBody {
	return &NameResponseBody{
		Session:session.NewStrippedSession(),
	}
}

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(ctx context.Context, request Request) (Response, error) {
	var buf bytes.Buffer

	log.Printf("Request: %+v", request)

	var requestBody *NameRequestBody // Unmarshal the body of the request
	err := json.Unmarshal([]byte(request.Body), &requestBody)

	session, err := whosthatpokemon.LoadGameSession(requestBody.SessionID)  // Load the new session

	if err != nil {
		return Response{StatusCode: 404}, err
	}

	// This does a basic check for profanity
	if goaway.IsProfane(requestBody.UserName) {
		return Response{StatusCode:422}, errors.New("profanity")
	}

	err = session.SetUserName(requestBody.UserName)

	if err != nil {
		return Response{StatusCode: 404}, err
	}

	body, err := json.Marshal(NewStartResponseBody(session))

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
			"Access-Control-Allow-Origin": "whosthatpokemon.xyz",
		},
	}
	return resp, nil
}

func main() {
	lambda.Start(Handler)
}
