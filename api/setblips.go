package main

import (
  "log"
  "encoding/json"
  "github.com/aws/aws-sdk-go/aws"
  "github.com/aws/aws-sdk-go/aws/session"
  "github.com/aws/aws-sdk-go/service/dynamodb"
  "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/satori/go.uuid"
)

// CreateBlip takes a user name and returns all blips as json info
func CreateBlip(username string, blip string) (retval string, err error) {

  log.Printf("creating blip for %s :: %s", username, blip)

  log.Printf("Creating a new blipid")
  blipid, err := uuid.NewV4()
  if err != nil {
    log.Printf("Error Creating uuid :: %v",err.Error())
    return "", err
  }

  log.Printf("Loading the blip object")
  item := Blip{
    BlipID: blipid.String(),
    UserName: username,
    Blip: blip,
  }

  av, err := dynamodbattribute.MarshalMap(item)
  if err != nil {
    log.Printf("Error marshalling item :: %v",err.Error())
    return "", err
  }

  log.Printf("Creating Session")
  sess, err := session.NewSession(&aws.Config{
      Region: aws.String("us-east-1")},
  )
  if err != nil {
    log.Printf("Error Creating Session :: %v",err.Error())
    return "", err
  }

  // Create DynamoDB client
  log.Printf("Creating DB Client")
  svc := dynamodb.New(sess)

  input := &dynamodb.PutItemInput{
    Item: av,
    TableName: aws.String("Blips"),
  }

  log.Printf("Putting item into table")
  _, err = svc.PutItem(input)
  if err != nil {
    log.Printf("Got error calling PutItem:")
    log.Printf(err.Error())
    return "", err
  }

  // return a result
  j, err := json.Marshal(item)
  if err != nil {
    log.Printf("Failed to convert to json, %v", err)
    return "", err
  }
  return string(j), nil

}

// UpdateBlip takes a user name and returns all blips as json info
func UpdateBlip(blipid string, username string, blip string) (retval string, err error) {

  log.Printf("updating blip id %s", blipid)

  log.Printf("Creating Session")
  sess, err := session.NewSession(&aws.Config{
      Region: aws.String("us-east-1")},
  )
  if err != nil {
    log.Printf("Error Creating Session :: %v",err.Error())
    return "", err
  }

  // Create DynamoDB client
  log.Printf("Creating DB Client")
  svc := dynamodb.New(sess)

  log.Printf("blipid %s", blipid)
  log.Printf("blip %s", blip)
  log.Printf("username %s", username)

  input := &dynamodb.UpdateItemInput{
    ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
      ":r": {
        S: aws.String(blip),
      },
    },
    TableName: aws.String("Blips"),
    Key: map[string]*dynamodb.AttributeValue{
      "blipid": {
        S: aws.String(blipid),
      },
    },
    ReturnValues:     aws.String("UPDATED_NEW"),
    UpdateExpression: aws.String("set blip = :r"),
  }
  log.Print(input)
  log.Printf("Updating item in table")
  _, err = svc.UpdateItem(input)
  if err != nil {
    log.Printf("Got error calling UpdateItem:")
    log.Printf(err.Error())
    return "", err
  }

  // use the item object to convert to json
  item := Blip{}
  item.BlipID = blipid
  item.UserName = username
  item.Blip = blip


  j, err := json.Marshal(item)
  if err != nil {
    log.Printf("Failed to convert to json, %v", err)
    return "", err
  }
  return string(j), nil
}
