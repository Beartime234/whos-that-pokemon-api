package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Response is of type APIGatewayProxyResponse since we're leveraging the
// AWS Lambda Proxy Request functionality (default behavior)
//
// https://serverless.com/framework/docs/providers/aws/events/apigateway/#lambda-proxy-integration
type Response events.APIGatewayProxyResponse

type Request struct {
	ID int `json:"id"`
}

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(request Request, ctx context.Context) (Response, error) {
	var buf bytes.Buffer

	body, err := json.Marshal(map[string]interface{}{
		"message": fmt.Sprintf("Your ID is %d", request.ID),
	})

	if err != nil {
		return Response{StatusCode: 404}, err
	}

	json.HTMLEscape(&buf, body)

	resp := Response{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            buf.String(),
		Headers: map[string]string{
			"Content-Type":           "application/json",
			"X-WTP-Func-Reply": "api-Handler",
		},
	}

	return resp, nil
}

func main() {
	lambda.Start(Handler)
}
