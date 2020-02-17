package main

import (
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type item struct {
	ChanelID  string
	Usernames []string
}

var (
	tableName = os.Getenv("TABLENAME")
)

func main() {
	lambda.Start(handleRequest)
}

func handleRequest(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	fmt.Println(request.QueryStringParameters)

	switch request.QueryStringParameters["eventType"] {
	// User connect management
	case "CONNECT":
		err := addConnection(request.QueryStringParameters["channelID"], request.QueryStringParameters["username"])
		if err != nil {
			log.Println(err)
		}
		return events.APIGatewayProxyResponse{
			StatusCode: 200,
			Body:       "Connexion réussie",
		}, err

	case "DISCONNECT":
		err := deleteConnection(request.QueryStringParameters["channelID"], request.QueryStringParameters["username"])
		if err != nil {
			log.Println(err)
		}
		return events.APIGatewayProxyResponse{
			StatusCode: 200,
			Body:       "déconnexion réussie",
		}, err

	default:
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       string(request.QueryStringParameters),
		}, nil
	}
}

// Connection Management
func addConnection(channelID, username string) error {
	updateItem(tableName, channelID, username)
	return nil
}

func deleteConnection(channelID, username string) error {
	return nil
}

/* DynamoDB interaction */
// Update an item into a DynamoDB table
func updateItem(tableName, channelID, username string) {
	// Create DynamoDB client
	mySession := session.Must(session.NewSession())

	// Create a DynamoDB client with additional configuration
	svc := dynamodb.New(mySession, aws.NewConfig().WithRegion("eu-west-1"))

	// Get the actual version of the item
	oldItem := readItem(tableName, channelID)

	// Append the actual Usernames list with the new username
	newUsernames := append(oldItem.Usernames, username)

	item := item{
		ChanelID:  channelID,
		Usernames: newUsernames,
	}

	// Marshal the Item into a readable format for DynamoDB
	av, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		fmt.Println("Got error marshalling new movie item:")
		fmt.Println(err.Error())
		return
	}

	// Create the PutItemInput object
	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(tableName),
	}

	// Put the item into the DynamoDB table
	_, err = svc.PutItem(input)
	if err != nil {
		fmt.Println("Got error calling PutItem:")
		fmt.Println(err.Error())
		return
	}
}

// Read an item into a DynamoDB table
func readItem(tableName, channelID string) item {
	// Create DynamoDB client
	mySession := session.Must(session.NewSession())

	// Create a DynamoDB client with additional configuration
	svc := dynamodb.New(mySession, aws.NewConfig().WithRegion("eu-west-1"))

	result, err := svc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"ChanelID": {
				S: aws.String(channelID),
			},
		},
	})
	if err != nil {
		fmt.Println(err.Error())
	}

	item := item{}

	err = dynamodbattribute.UnmarshalMap(result.Item, &item)
	if err != nil {
		panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
	}

	return item
}

//
// Refaire une fonction delete en se basant sur update en enlevant le username de l'utilisateur voulant se déconnecter
// deleteItem permet de détruire un channel
//

// delete an item into a DynamoDB table
func deleteItem(tableName, channelID string) {
	// Create DynamoDB client
	mySession := session.Must(session.NewSession())

	// Create a DynamoDB client with additional configuration
	svc := dynamodb.New(mySession, aws.NewConfig().WithRegion("eu-west-1"))

	// Create the DeleteItemInput object
	input := &dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"ChanelID": {
				S: aws.String(channelID),
			},
		},
		TableName: aws.String(tableName),
	}

	// Delete the item into the DynamoDB table
	_, err := svc.DeleteItem(input)
	if err != nil {
		fmt.Println("Got error calling DeleteItem")
		fmt.Println(err.Error())
		return
	}
}
