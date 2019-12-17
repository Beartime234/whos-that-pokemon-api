package main

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"os"
)

// Response is of type APIGatewayProxyResponse since we're leveraging the
// AWS Lambda Proxy Request functionality (default behavior)
//
// https://serverless.com/framework/docs/providers/aws/events/apigateway/#lambda-proxy-integration
type Response events.APIGatewayProxyResponse

type Request struct {

}

var GalleryTableName = os.Getenv("GALLERY_TABLE_NAME")
const GalleryTableHashKey = "PokedexID"
var SessionTableName = os.Getenv("SESSION_TABLE_NAME")
const SessionTableHashKey = "SessionID"

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(ctx context.Context, request *Request) (Response, error) {
	var buf bytes.Buffer

	session := NewGameSession()  // Create a new session

	err := session.save() // Save the session in the database so we can get it later

	if err != nil {
		return Response{StatusCode: 404}, err
	}

	strippedSession := session.NewStrippedSession()

	body, err := json.Marshal(strippedSession)

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
