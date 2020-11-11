package repo

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/rociosantos/gokit-example/service"
	"github.com/sirupsen/logrus"
)

type repo struct {
	d      *DynamoClient
	logger *logrus.Logger
}

func NewRepo(db *DynamoClient, logger *logrus.Logger) service.Repository {
	return &repo{
		d:      db,
		logger: logger,
	}
}

type UserItem struct {
	Id        string `dynamodbav:"UserId"`
	Email     string `dynamodbav:"Email"`
	Password  string `dynamodbav:"Password"`
	CreatedAt int64 `dynamodbav:"CreatedAt"`
	UpdatedAt int64 `dynamodbav:"UpdatedAt"`
}

func (repo *repo) CreateUser(ctx context.Context, user service.User) error {

	item := UserItem{
		user.ID,
		user.Email,
		user.Password,
		time.Now().Unix(),
		time.Now().Unix(),
	}
	attV, err := dynamodbattribute.MarshalMap(item)
	if err != nil{
		return err
	}

	_, err = repo.d.dbClient.PutItem(&dynamodb.PutItemInput{
		ConditionExpression: aws.String("attribute_not_exists(id) AND" +
			" attribute_not_exists(email)"),
		Item:                attV,
		TableName:           aws.String(repo.d.tableNames["users"]),
	})

	repo.logger.Debug("User created with id ", user.ID)

	if err != nil {
		return err
	}
	return nil
}

func (repo *repo) GetUser(ctx context.Context, id string) (string, error) {
	result, err := repo.d.dbClient.Query(&dynamodb.QueryInput{
		KeyConditionExpression: aws.String("UserId = :uid"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":uid": {S: aws.String(id)},
		},
		TableName: aws.String(repo.d.tableNames["users"]),
	})
	if err != nil {
		return "", err
	}

	item := service.User{}
	err = dynamodbattribute.UnmarshalMap(result.Items[0], &item)
	if err != nil {
		return "", err
	}

	return item.Email, nil
}
