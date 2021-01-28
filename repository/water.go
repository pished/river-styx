package repository

import (
	"fmt"
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"github.com/google/uuid"
	"github.com/pished/river-styx/model"
	"github.com/spf13/viper"
)

func ListWater(conn *dynamodb.DynamoDB) ([]model.WaterForm, error) {
	filt := expression.Name("Category").Equal(expression.Value("Water"))
	proj := expression.NamesList(
		expression.Name("ItemId"),
		expression.Name("Type"),
		expression.Name("Brand"),
		expression.Name("Category"),
		expression.Name("Flavored"),
		expression.Name("Flavoring"),
		expression.Name("Quality"),
	)
	expr, err := expression.NewBuilder().WithFilter(filt).WithProjection(proj).Build()
	if err != nil {
		return nil, err
	}

	params := &dynamodb.ScanInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
		ProjectionExpression:      expr.Projection(),
		TableName:                 aws.String(viper.GetString("db.table")),
	}
	result, err := conn.Scan(params)
	if err != nil {
		return nil, err
	}

	waters := make([]model.WaterForm, 0)
	for _, i := range result.Items {
		item := model.WaterForm{}
		err = dynamodbattribute.UnmarshalMap(i, &item)
		fmt.Println(item)
		waters = append(waters, item)
	}

	return waters, nil
}

func AddWater(conn *dynamodb.DynamoDB, body *model.WaterModel) error {
	tempo := model.WaterModel{
		Type:      body.Type,
		Brand:     body.Brand,
		Category:  body.Category,
		Flavored:  body.Flavored,
		Flavoring: body.Flavoring,
		Quality:   body.Quality,
	}
	fmt.Println(tempo)
	body.ItemId = uuid.NewString()
	av, err := dynamodbattribute.MarshalMap(body)
	if err != nil {
		fmt.Println("Error Marshalling map")
		fmt.Println(err)
	}
	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(viper.GetString("db.table")),
	}
	_, noSuccess := conn.PutItem(input)

	return noSuccess
}

func GetWater(conn *dynamodb.DynamoDB, itemId string) ([]model.WaterForm, error) {
	filt := expression.Name("ItemId").Equal(expression.Value(itemId))
	proj := expression.NamesList(
		expression.Name("ItemId"),
		expression.Name("Type"),
		expression.Name("Brand"),
		expression.Name("Category"),
		expression.Name("Flavored"),
		expression.Name("Flavoring"),
		expression.Name("Quality"),
	)
	expr, err := expression.NewBuilder().WithFilter(filt).WithProjection(proj).Build()

	if err != nil {
		return nil, err
	}

	params := &dynamodb.ScanInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
		ProjectionExpression:      expr.Projection(),
		TableName:                 aws.String(viper.GetString("db.table")),
	}
	result, err := conn.Scan(params)
	if err != nil {
		return nil, err
	}

	var waters model.Waters
	for _, i := range result.Items {
		item := model.WaterForm{}
		err = dynamodbattribute.UnmarshalMap(i, &item)
		fmt.Println(item)
		waters = append(waters, item)
	}

	return waters, nil
}

func UpdateWater(conn *dynamodb.DynamoDB, body *model.WaterModel, itemId string) error {
	rate := strconv.Itoa(body.Quality)
	info := map[string]string{
		"ItemId": itemId,
	}

	key, err := dynamodbattribute.MarshalMap(info)
	if err != nil {
		return err
	}

	input := &dynamodb.UpdateItemInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":r": {
				N: aws.String(rate),
			},
		},
		TableName:        aws.String(viper.GetString("db.table")),
		Key:              key,
		ReturnValues:     aws.String("UPDATED_NEW"),
		UpdateExpression: aws.String("set Quality = :r"),
	}
	_, noSuccess := conn.UpdateItem(input)
	return noSuccess
}

func DeleteWater(conn *dynamodb.DynamoDB, itemId string) error {
	input := &dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"ItemId": {
				S: aws.String(itemId),
			},
		},
		TableName: aws.String(viper.GetString("db.table")),
	}
	_, err := conn.DeleteItem(input)

	return err
}
