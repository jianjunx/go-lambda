package dao

import (
	"go-zone/config"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

var db *dynamo.DB

func getTable(name string) dynamo.Table {
	return db.Table(name)
}

func DynamoInit() {
	sess := session.Must(session.NewSession())
	region := config.GetConfig().Dynamo.Region
	db = dynamo.New(sess, &aws.Config{Region: aws.String(region)})
}

func getPostTable() dynamo.Table {
	return getTable(config.GetConfig().PostTableName)
}

func getBookTable() dynamo.Table {
	return getTable(config.GetConfig().BookTableName)
}
