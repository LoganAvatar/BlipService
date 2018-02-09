FROM golang:latest

ARG PACKAGE="github.com/loganavatar/BlipService"
WORKDIR /go/src/$PACKAGE
RUN \
  go get -d "github.com/aws/aws-lambda-go/events" && \
  go get -d "github.com/aws/aws-lambda-go/lambda" && \
  go get -d "github.com/aws/aws-sdk-go/aws" && \
  go get -d "github.com/aws/aws-sdk-go/aws/session" && \
  go get -d "github.com/aws/aws-sdk-go/service/dynamodb" && \
  go get -d "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute" && \
  go get -d "github.com/aws/aws-sdk-go/service/dynamodb/expression" && \
  go get -d "github.com/satori/go.uuid"

COPY . .

RUN \
  cd api && \
  go get -u github.com/golang/lint/golint && \
  go get -d -v ./... && \
  golint -set_exit_status && \
  go vet . && \
  go build -v ./...
