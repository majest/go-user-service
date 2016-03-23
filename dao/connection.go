package dao

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

var conn *dynamodb.DynamoDB
var config *aws.Config

// ConfigureDb sets the config object so that it can be used when connection is needed
func ConfigureDb(hostname, region string, port int) {
	dbHostnameWithPort := fmt.Sprintf("http://%s:%d", hostname, port)
	config = &aws.Config{
		Region:   &region,
		Endpoint: &dbHostnameWithPort,
	}
}

// GetConn returns dynamodb connection
func GetConn() *dynamodb.DynamoDB {

	if conn == nil {
		conn = dynamodb.New(session.New(config))
	}

	return conn
}
