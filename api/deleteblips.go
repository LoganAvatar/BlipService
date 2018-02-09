package main

import (
    "log"

    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/dynamodb"
)


// DeleteBlip takes a user name and returns all blips as json info
func DeleteBlip(blipid string) (err error) {

  log.Printf("deleting blip id %s", blipid)
  log.Printf("Creating Session")
  sess, err := session.NewSession(&aws.Config{
    Region: aws.String("us-east-1")},
  )
  if err != nil {
    log.Printf("Error Creating Session :: %v",err.Error())
    return err
  }

  // Create DynamoDB client
  log.Printf("Creating DB Client")
  svc := dynamodb.New(sess)

  input := &dynamodb.DeleteItemInput{
    Key: map[string]*dynamodb.AttributeValue{
      "blipid": {
        S: aws.String(blipid),
      },
    },
    TableName: aws.String("Blips"),
  }

  _, err = svc.DeleteItem(input)
  if err != nil {
    log.Printf("Got error calling DeleteItem")
    log.Printf(err.Error())
    return err
  }

  return nil
}
