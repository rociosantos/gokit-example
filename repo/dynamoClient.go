package repo

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/sirupsen/logrus"
	"github.com/aws/aws-sdk-go/aws/session"
)

// DynamoConnection - Functions to handle the database
type DynamoConnection interface {
	ListTables(*dynamodb.ListTablesInput) (*dynamodb.ListTablesOutput, error)
	PutItem(*dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error)
	GetItem(*dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error)
	Query(input *dynamodb.QueryInput) (*dynamodb.QueryOutput, error)
}

type DynamoClient struct {
	tableNames map[string]string
	dbClient DynamoConnection
	logger *logrus.Logger
}

func NewDynamoClient(tables map[string]string, s *session.Session, logger *logrus.Logger) *DynamoClient {
	return &DynamoClient{
		tableNames: tables,
		dbClient: dynamodb.New(s),
		logger: logger,
	}
}

