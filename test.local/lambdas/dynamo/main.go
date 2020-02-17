package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type Item struct {
	ChanelID  string
	Usernames []string
}

type ItemRead struct {
	ChanelID  string
	Usernames []string
}

func main() {
	lambda.Start(handleRequest)
}

func handleRequest() {
	//listTables()

	/*
		createItem("testChanel1")
		createItem("testChanel3")
	*/

	createItem("testChanel2")
	updateItem()

	//readItem()

	//deleteItem()
}

func listTables() {
	mySession := session.Must(session.NewSession())

	// Create a DynamoDB client with additional configuration
	svc := dynamodb.New(mySession, aws.NewConfig().WithRegion("eu-west-1"))

	input := &dynamodb.ListTablesInput{}

	result, err := svc.ListTables(input)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(result)
}

func createItem(channelID string) {
	// Create DynamoDB client
	mySession := session.Must(session.NewSession())

	// Create a DynamoDB client with additional configuration
	svc := dynamodb.New(mySession, aws.NewConfig().WithRegion("eu-west-1"))

	item := Item{
		ChanelID: channelID,
		Usernames: []string{
			"admin2",
		},
	}

	av, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		fmt.Println("Got error marshalling new movie item:")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	tableName := "test-chat"

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(tableName),
	}

	_, err = svc.PutItem(input)
	if err != nil {
		fmt.Println("Got error calling PutItem:")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println("create new chanel")
}

func updateItem() {
	// Create DynamoDB client
	mySession := session.Must(session.NewSession())

	// Create a DynamoDB client with additional configuration
	svc := dynamodb.New(mySession, aws.NewConfig().WithRegion("eu-west-1"))

	oldItem := readItem()

	newUsernames := append(oldItem.Usernames, "valentin")

	item := Item{
		ChanelID:  "test",
		Usernames: newUsernames,
	}

	av, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		fmt.Println("Got error marshalling new movie item:")
		fmt.Println(err.Error())
	}

	tableName := "test-chat"

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(tableName),
	}

	_, err = svc.PutItem(input)
	if err != nil {
		fmt.Println("Got error calling PutItem:")
		fmt.Println(err.Error())
	}

	fmt.Println("update chanel")
}

/*
func updateItem(channelID string) {
	// Create DynamoDB client
	mySession := session.Must(session.NewSession())

	// Create a DynamoDB client with additional configuration
	svc := dynamodb.New(mySession, aws.NewConfig().WithRegion("eu-west-1"))

	// Get the actual item
	oldItem := readItem()

	// Item identifies the item in the table
	newUsernames := append(oldItem.Usernames, "valentin")

    tableName := "test"

	input := &dynamodb.UpdateItemInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":u": {
				L: {
					"kevin",
					"valentin",
				},
			},
		},
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"ChanelID": {
				N: aws.String("test"),
			},
		},
		ReturnValues:     aws.String("UPDATED_NEW"),
		UpdateExpression: aws.String("set Usernames = :u"),
	}

	_, err := svc.UpdateItem(input)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("update item in db")
}
*/

func readItem() ItemRead {
	// Create DynamoDB client
	mySession := session.Must(session.NewSession())

	// Create a DynamoDB client with additional configuration
	svc := dynamodb.New(mySession, aws.NewConfig().WithRegion("eu-west-1"))

	result, err := svc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String("test-chat"),
		Key: map[string]*dynamodb.AttributeValue{
			"ChanelID": {
				S: aws.String("test"),
			},
		},
	})

	fmt.Println(result)

	if err != nil {
		fmt.Println(err.Error())
	}

	item := ItemRead{}

	err = dynamodbattribute.UnmarshalMap(result.Item, &item)

	if err != nil {
		panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
	}

	fmt.Println("Found item:")
	fmt.Println("chanelID:  ", item.ChanelID)
	fmt.Println("Usernames: ", item.Usernames)

	return item
}

func deleteItem() {
	// Create DynamoDB client
	mySession := session.Must(session.NewSession())

	// Create a DynamoDB client with additional configuration
	svc := dynamodb.New(mySession, aws.NewConfig().WithRegion("eu-west-1"))

	input := &dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"ChanelID": {
				S: aws.String("testChanel2"),
			},
		},
		TableName: aws.String("test-chat"),
	}

	_, err := svc.DeleteItem(input)
	if err != nil {
		fmt.Println("Got error calling DeleteItem")
		fmt.Println(err.Error())
		return
	}

	fmt.Println("delete chat testChanel12")
}
