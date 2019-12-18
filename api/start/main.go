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

}

type StartResponseBody struct {
	Session *whosthatpokemon.GameSession
}

func NewStartResponseBody(session *whosthatpokemon.GameSession) *StartResponseBody {
	return &StartResponseBody{Session: session}
}

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(ctx context.Context, request *Request) (Response, error) {
	var buf bytes.Buffer

	log.Printf("Request: %+v", request)

	session, err := whosthatpokemon.NewGameSession()  // Create a new session

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
		},
	}
	return resp, nil
}

func main() {
	lambda.Start(Handler)
}
