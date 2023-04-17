package handlers

import (
	"com.mormorbump/apigateway.dynamodb.websockets.golang/domain/models"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"log"
	"strconv"
	"time"
)

type DBHandler struct {
	client     *dynamodb.Client
	primaryKey string
	sortKey    string
	tableName  string
	gsiName    string
	ttlKey     string
	attrKey1   string
}

const Expire = 10 * time.Hour

func NewDBHandler(cfg aws.Config, pkey string, skey string, tblName string, gsiName string, ttlKey string, attrKey1 string) *DBHandler {
	client := dynamodb.NewFromConfig(cfg)
	return &DBHandler{
		client:     client,
		primaryKey: pkey,
		sortKey:    skey,
		tableName:  tblName,
		gsiName:    gsiName,
		ttlKey:     ttlKey,
		attrKey1:   attrKey1,
	}
}

func (db *DBHandler) DBPutItem(primaryKeyVal string, sortKeyVal string, attrKey1Val string) error {
	_, err := db.client.PutItem(context.TODO(), db.setPutItemInput(primaryKeyVal, sortKeyVal, attrKey1Val))
	return err
}

func (db *DBHandler) DBGetItem(primaryKeyVal string, sortKeyVal string, out interface{}) error {
	getItemOutPut, err := db.client.GetItem(context.TODO(), db.setGetItemInput(primaryKeyVal, sortKeyVal))
	if getItemOutPut == nil {
		return nil
	}
	err = attributevalue.UnmarshalMap(getItemOutPut.Item, &out)
	return err
}

func (db *DBHandler) DBDeleteItem(sortKeyVal string) error {
	_, err := db.client.DeleteItem(context.TODO(), db.setDeleteItemInput(sortKeyVal))
	return err
}

func (db *DBHandler) DBQuery(primaryKey string, out interface{}) error {
	queryOutPut, err := db.client.Query(context.TODO(), db.setQueryInput(primaryKey))
	if err != nil {
		return err
	}
	err = attributevalue.UnmarshalListOfMaps(queryOutPut.Items, &out)
	if err != nil {
		return err
	}
	return nil
}

func (db *DBHandler) setPutItemInput(primaryKeyVal string, sortKeyVal string, attrKey1Val string) *dynamodb.PutItemInput {
	return &dynamodb.PutItemInput{
		TableName: aws.String(db.tableName),
		Item: map[string]types.AttributeValue{
			db.primaryKey: &types.AttributeValueMemberS{Value: primaryKeyVal},
			db.sortKey:    &types.AttributeValueMemberS{Value: sortKeyVal},
			db.ttlKey:     &types.AttributeValueMemberN{Value: strconv.FormatInt(setTtl(Expire), 10)},
			db.attrKey1:   &types.AttributeValueMemberS{Value: attrKey1Val},
		},
	}
}

func (db *DBHandler) setGetItemInput(primaryKeyVal string, sortKeyVal string) *dynamodb.GetItemInput {
	return &dynamodb.GetItemInput{
		TableName: aws.String(db.tableName),
		Key: map[string]types.AttributeValue{
			db.primaryKey: &types.AttributeValueMemberS{Value: primaryKeyVal},
			db.sortKey:    &types.AttributeValueMemberS{Value: sortKeyVal},
		},
		ConsistentRead: aws.Bool(true),
	}
}

func (db *DBHandler) setDeleteItemInput(GSIHashKeyVal string) *dynamodb.DeleteItemInput {

	out, err := db.getIndexQueryInput(GSIHashKeyVal)
	if err != nil {
		log.Println(err)
	}
	return &dynamodb.DeleteItemInput{
		Key: map[string]types.AttributeValue{
			db.primaryKey: &types.AttributeValueMemberS{Value: out[0].RoomID},
			db.sortKey:    &types.AttributeValueMemberS{Value: out[0].UserID},
		},
		TableName: aws.String(db.tableName),
	}
}

func (db *DBHandler) getIndexQueryInput(GSIHashKeyVal string) (out models.Connections, err error) {
	queryOutPut, err := db.client.Query(context.TODO(), db.setIndexQueryInput(GSIHashKeyVal))
	if err != nil {
		log.Println(err)
	}
	out = models.Connections{}
	err = attributevalue.UnmarshalListOfMaps(queryOutPut.Items, &out)
	return out, err
}

func (db *DBHandler) setIndexQueryInput(GSIHashKeyVal string) *dynamodb.QueryInput {
	keyCond := expression.Key(db.attrKey1).Equal(expression.Value(GSIHashKeyVal))
	proj := expression.NamesList(expression.Name(db.sortKey), expression.Name(db.primaryKey), expression.Name(db.attrKey1))
	expr, err := expression.NewBuilder().
		WithKeyCondition(keyCond).
		WithProjection(proj).
		Build()
	if err != nil {
		fmt.Println(err)
	}

	return &dynamodb.QueryInput{
		TableName:                 aws.String(db.tableName),
		IndexName:                 aws.String(db.gsiName),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
		KeyConditionExpression:    expr.KeyCondition(),
		ProjectionExpression:      expr.Projection(),
	}
}

func (db *DBHandler) setQueryInput(primaryKeyVal string) *dynamodb.QueryInput {

	keyCond := expression.Key(db.primaryKey).Equal(expression.Value(primaryKeyVal))
	proj := expression.NamesList(expression.Name("connectionId"))
	expr, err := expression.NewBuilder().
		WithKeyCondition(keyCond).
		WithProjection(proj).
		Build()
	if err != nil {
		fmt.Println(err)
	}

	return &dynamodb.QueryInput{
		TableName:                 aws.String(db.tableName),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
		KeyConditionExpression:    expr.KeyCondition(),
		ProjectionExpression:      expr.Projection(),
	}
}

func setTtl(expire time.Duration) int64 {
	now := time.Now()
	return now.Add(expire).Unix()
}
