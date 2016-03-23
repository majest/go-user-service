package dao

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type Tables struct {
	Svc *dynamodb.DynamoDB
}

func NewTable() *Tables {
	return &Tables{Svc: GetConn()}
}

func (t *Tables) CreateUsersTable() (*dynamodb.CreateTableOutput, error) {

	params := &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{ // Required
			{
				AttributeName: aws.String("Id"),
				AttributeType: aws.String("S"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("Id"),
				KeyType:       aws.String("HASH"),
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{ // Required
			ReadCapacityUnits:  aws.Int64(1),
			WriteCapacityUnits: aws.Int64(1),
		},
		TableName: aws.String("users"),
	}
	resp, err := t.Svc.CreateTable(params)

	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(resp)
	return resp, err
}

func (t *Tables) ListAll() {

	params := &dynamodb.ListTablesInput{
		//ExclusiveStartTableName: aws.String("TableName"),
		Limit: aws.Int64(10),
	}
	resp, err := t.Svc.ListTables(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}
