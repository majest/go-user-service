package dao

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/majest/user-service/server"
	"github.com/satori/go.uuid"
)

var fields []string = []string{"Id", "FirstName", "LastName", "Email", "Address", "Street", "PostCode", "City", "Country", "Phone"}

type User struct {
	Svc       *dynamodb.DynamoDB
	tableName *string
}

func NewUserDao() *User {
	return &User{Svc: GetConn(), tableName: aws.String("users")}
}

func (u *User) Save(user *server.User) (string, error) {

	if user.Id == "" {
		user.Id = uuid.NewV4().String()
	}

	output, err := u.Svc.PutItem(&dynamodb.PutItemInput{
		Item:      getItems(user),
		TableName: u.tableName, // Required
	})

	fmt.Println(output)
	return user.Id, err
}

func (u *User) FindOne(user *server.UserSearch) (*server.User, error) {

	output, err := u.Svc.GetItem(&dynamodb.GetItemInput{
		Key:       getItems(user),
		TableName: u.tableName,
	})

	if output.Item != nil {

		su := &server.User{}

		if output.Item["Id"] != nil {
			su.Id = *output.Item["Id"].S
		}

		if output.Item["FirstName"] != nil {
			su.FirstName = *output.Item["FirstName"].S
		}

		if output.Item["LastName"] != nil {
			su.LastName = *output.Item["LastName"].S
		}

		if output.Item["Email"] != nil {
			su.Email = *output.Item["Email"].S
		}

		if output.Item["Address"] != nil {
			su.Address = *output.Item["Address"].S
		}

		if output.Item["Street"] != nil {
			su.Street = *output.Item["Street"].S
		}

		if output.Item["PostCode"] != nil {
			su.PostCode = *output.Item["PostCode"].S
		}

		if output.Item["City"] != nil {
			su.City = *output.Item["City"].S
		}

		if output.Item["Country"] != nil {
			su.Country = *output.Item["Country"].S
		}

		if output.Item["Phone"] != nil {
			su.Phone = *output.Item["Phone"].S
		}

		return su, err
	}

	return nil, err

}
