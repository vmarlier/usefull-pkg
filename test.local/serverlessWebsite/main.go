package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(handleRequest)
}

func handleRequest(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       fmt.Sprintf("%s, %s", "Réponse", request.Body),
	}, nil
}
