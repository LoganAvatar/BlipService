package main

import (
  "errors"
  "log"
  "strings"
  "encoding/json"
  "github.com/aws/aws-lambda-go/events"
  "github.com/aws/aws-lambda-go/lambda"
)

var (
  // ErrBadRequest is thrown when a request isn't right
  ErrBadRequest = errors.New("Bad Request, try better next time")
)

// Blip is the representation of the db item
type Blip struct {
    Blip string`json:"blip"`
    BlipID string`json:"blipid"`
    UserName string`json:"username"`
}

// Handler is your Lambda function handler
// It uses Amazon API Gateway request/responses provided by the aws-lambda-go/events package,
// However you could use other event sources (S3, Kinesis etc), or JSON-decoded primitive types such as 'string'.
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

  // Print out the request info
  log.Printf("Processing Lambda request %s against %s\n", request.RequestContext.RequestID, request.Resource)

  // log the request object
  e, err := json.Marshal(request)
  if err != nil {
    log.Print(err)
    return events.APIGatewayProxyResponse{
      Body: "error " + request.Body,
      StatusCode: 400,
    }, err
  }
  log.Printf(string(e))

  // switch between each supported Function, doing stuff
  if strings.Contains(request.Path,"/get/user/") {
    userid := strings.Split(request.Path, "/")[3]
    retval, err := GetBlips(userid)
    if err != nil {
      log.Print(err)
      return events.APIGatewayProxyResponse{
        Body: err.Error(),
        StatusCode: 400,
      }, err
    }

    // return stuff
    return events.APIGatewayProxyResponse{
      Body: retval,
      StatusCode: 200,
    }, nil

  } else if strings.Contains(request.Path,"/get/") {
    // the id should be the 3rd item
    blipid := strings.Split(request.Path, "/")[2]
    retval, err := GetBlipItem(blipid)
    if err != nil {
      log.Print(err)
      return events.APIGatewayProxyResponse{
        Body: err.Error(),
        StatusCode: 400,
      }, err
    }

    // return stuff
    return events.APIGatewayProxyResponse{
      Body: retval,
      StatusCode: 200,
    }, nil

  } else if strings.Contains(request.Path,"/delete") {

    // Parse the body
    _, err := json.Marshal(request.Body)
    if err != nil {
      log.Printf("Failed to parse json, %v", err)
      return events.APIGatewayProxyResponse{
        Body: "Failed processing",
        StatusCode: 400,
      }, err
    }

    // create a blip object and load it
    blip := Blip{}
    errx := json.Unmarshal([]byte(request.Body), &blip)
    if errx != nil {
      log.Print(err)
      return events.APIGatewayProxyResponse{
        Body: err.Error(),
        StatusCode: 400,
      }, errx
    }

    // delete the blip
    erry := DeleteBlip(blip.BlipID)
    if erry != nil {
      log.Print(err)
      return events.APIGatewayProxyResponse{
        Body: err.Error(),
        StatusCode: 400,
      }, erry
    }

    // return stuff
    return events.APIGatewayProxyResponse{
      Body: "Blip Deleted",
      StatusCode: 200,
    }, nil

  } else if strings.Contains(request.Path,"/set") {

    // Parse the body
    _, err := json.Marshal(request.Body)
    if err != nil {
      log.Printf("Failed to parse json, %v", err)
      return events.APIGatewayProxyResponse{
        Body: "Failed processing",
        StatusCode: 400,
      }, err
    }

    // create a blip object and load it
    blip := Blip{}
    errx := json.Unmarshal([]byte(request.Body), &blip)
    if errx != nil {
      log.Print(err)
      return events.APIGatewayProxyResponse{
        Body: err.Error(),
        StatusCode: 400,
      }, errx
    }

    // check to see if a blipid is in the json
    if blip.BlipID != "" {
      // update the blip
      retval, err := UpdateBlip(blip.BlipID, blip.UserName, blip.Blip)
      if err != nil {
        log.Print(err)
        return events.APIGatewayProxyResponse{
          Body: err.Error(),
          StatusCode: 400,
        }, err
      }

      // return stuff
      return events.APIGatewayProxyResponse{
        Body: retval,
        StatusCode: 200,
      }, nil

    }

    // create the blip
    retval, err := CreateBlip(blip.UserName, blip.Blip)
    if err != nil {
      log.Print(err)
      return events.APIGatewayProxyResponse{
        Body: err.Error(),
        StatusCode: 400,
      }, err
    }

    // return stuff
    return events.APIGatewayProxyResponse{
      Body: retval,
      StatusCode: 200,
    }, nil

  }

  // return error
  return events.APIGatewayProxyResponse{
    Body: "Bad Request",
    StatusCode: 400,
  }, ErrBadRequest

}

// main is the loader for the lambda handler
func main() {
  lambda.Start(Handler)
}
