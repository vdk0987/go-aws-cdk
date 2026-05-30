package database

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

const (
	TABLE_NAME = "userTable"
)

type DynamoDBClient struct {
	databasestore *dynamodb.Client
}

func NewDynamoDBClient() (*DynamoDBClient, error) {
	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		panic(err)
	}
	db := dynamodb.NewFromConfig(cfg)

	return &DynamoDBClient{
		databasestore: db,
	}, nil
}

func (u *DynamoDBClient) UserValidation(username string) (bool, error) {
	result, err := u.databasestore.GetItem(
		context.Background(),
		&dynamodb.GetItemInput{
			TableName: aws.String(TABLE_NAME),
			Key: map[string]types.AttributeValue{
				username: &types.AttributeValueMemberS{
					Value: username,
				},
			},
		},
	)
	if err != nil {
		return true, err
	}
	if result.Item == nil {
		return false, nil
	}
	return true, nil
}
