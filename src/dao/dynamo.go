package dao

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

var svc *dynamodb.Client

func GetDB() *dynamodb.Client {
	return svc
}

func DynamoInit() {
	// Using the SDK's default configuration, loading additional config
	// and credentials values from the environment variables, shared
	// credentials, and shared configuration files
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("ap-northeast-1"))
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}
	// Using the Config value, create the DynamoDB client
	svc = dynamodb.NewFromConfig(cfg)

}

func GetTableNames() ([]string, error) {
	// Build the request with its input parameters
	resp, err := svc.ListTables(context.TODO(), &dynamodb.ListTablesInput{
		Limit: aws.Int32(5),
	})
	if err != nil {
		log.Fatalf("failed to list tables, %v", err)
		return nil, err
	}

	fmt.Println("Tables:")
	for _, tableName := range resp.TableNames {
		fmt.Println(tableName)
	}
	return resp.TableNames, nil
}
