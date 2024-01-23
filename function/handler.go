package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Handler is the Lambda function handler
func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	name := request.QueryStringParameters["name"]
	if name == "" {
		name = "Anonymous"
	}

	response := fmt.Sprintf("Hello, %s! Your Lambda function is working.", name)

	return events.APIGatewayProxyResponse{
		Body:       response,
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(Handler)
}
