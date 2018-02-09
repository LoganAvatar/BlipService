package main

import (
  "errors"
  "log"
  "encoding/json"
  "github.com/aws/aws-sdk-go/aws"
  "github.com/aws/aws-sdk-go/aws/session"
  "github.com/aws/aws-sdk-go/service/dynamodb"
  "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
  "github.com/aws/aws-sdk-go/service/dynamodb/expression"
)

// GetBlips takes a user name and returns all blips as json info
func GetBlips(username string) (retval string, err error) {

  log.Printf("Pulling user name %s", username)
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

  // set a filter
  log.Printf("Creating Scan filter expression")
  filt := expression.Name("username").Equal(expression.Value(username))

  // Get back the title, year, and rating
  log.Printf("Setting return dataset")
  proj := expression.NamesList(expression.Name("blip"), expression.Name("blipid"), expression.Name("username"))

  // get the expression with filter and projection
  log.Printf("Build the expression with filter and projection")
  expr, err := expression.NewBuilder().WithFilter(filt).WithProjection(proj).Build()
  if err != nil {
    log.Printf(err.Error())
    return "", err
  }

  log.Print(expr)

  params := &dynamodb.ScanInput{
      ExpressionAttributeNames:  expr.Names(),
      ExpressionAttributeValues: expr.Values(),
      FilterExpression:          expr.Filter(),
      ProjectionExpression:      expr.Projection(),
      TableName:                 aws.String("Blips"),
  }

  // Make the DynamoDB Query API call
  log.Printf("Scanning Table")
  result, err := svc.Scan(params)
  if err != nil {
    log.Printf(err.Error())
    return "", err
  }

  // return 0 results error
  if len(result.Items) == 0 {
    return "", errors.New("No Results returned")
  }

  // create the array to add blips to
  var blipslice []Blip
  for _, i := range result.Items {
    item := Blip{}
    // unmarshal to struct
    err = dynamodbattribute.UnmarshalMap(i, &item)
    if err != nil {
        log.Printf("Got error unmarshalling:")
        log.Printf(err.Error())
        return "", err
    }

    blipslice = append(blipslice, item)

  }

  // return a result
  j, err := json.Marshal(blipslice)
  if err != nil {
    log.Printf("Failed to convert to json, %v", err)
    return "", err
  }
  return string(j), nil

}

// GetBlipItem takes a user name and returns all blips as json info
func GetBlipItem(blipid string) (retval string, err error) {

  log.Printf("Pulling blip id %s", blipid)
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

  result, err := svc.GetItem(&dynamodb.GetItemInput{
    TableName: aws.String("Blips"),
    Key: map[string]*dynamodb.AttributeValue{
      "blipid": {
        S: aws.String(blipid),
      },
    },
  })
  if err != nil {
    log.Printf(err.Error())
    return "", err
  }
  log.Print(result)
  // Create a blank blip item and then load it up
  item := Blip{}
  err = dynamodbattribute.UnmarshalMap(result.Item, &item)
  if err != nil {
    log.Printf("Failed to unmarshal Record, %v", err)
    return "", err
  }

  if item.BlipID == "" {
    log.Printf("Could not find %s", blipid)
    return "", errors.New("blip id not found")
  }

  j, err := json.Marshal(item)
  if err != nil {
    log.Printf("Failed to convert to json, %v", err)
    return "", err
  }
  return string(j), nil
}
