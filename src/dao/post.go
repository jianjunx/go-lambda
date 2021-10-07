package dao

import (
	"context"
	"gin-template/src/model"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

var YUTABLENAME = "YQ_POST_TABLE"

func AddPostItemData(row *model.PostAttributes) (*dynamodb.PutItemOutput, error) {
	item := make(map[string]types.AttributeValue)
	item["title"] = row
	return GetDB().PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String("YQ_POST_TABLE"),
		Item:      item,
	})
}
