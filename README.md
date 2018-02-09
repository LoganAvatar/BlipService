# BlipService

### Requirements
The PM has sent me the following customer use cases
- Store Blips: simple text string.
- Retrieve own Blips.
- Retrieve other's Blips.
- Delete own Blips.

My tasks are:
- Build the Blip Service meeting the customer use cases
- Deploy into AWS
- Use metrics and monitoring to understand the Service

## Components

### Source Code
- Github Repo

### Lambda functions
- Blips-GetBlips
- Blips-StoreBlip
- Blips-DeleteBlip

### DynamoDB
#### Tables
- Blips

### IAM
- Policies
  - Blips-DB-RO
  - Blips-DB-RW
  - Blips-S3-
- Roles
  - Blips-BlipLambdaGet
  - Blips-BlipLambdaSet

### Route53

### S3
  Bucket - BlipsService

### Security Groups
### Route Tables
### Nat Gateway
### API Gateway
### CodePipeline
### CodeBuild

## Building
Test in docker build by running the following: `docker build --tag blipbuilder .`


## Executing
create blips
```bash
echo '{"blip":"Blipping what","username":"mike@mkdavies.com"}' > payload.txt
curl -XPOST -H "Content-Type: application/json" -d @payload.txt https://wkcfzdti0i.execute-api.us-east-1.amazonaws.com/prod/set
echo '{"blip":"This is my Second Blip","username":"mike@mkdavies.com"}' > payload.txt
curl -XPOST -H "Content-Type: application/json" -d @payload.txt https://wkcfzdti0i.execute-api.us-east-1.amazonaws.com/prod/set
echo '{"blip":"This is Rachels First Blip","username":"rachel@mkdavies.com"}' > payload.txt
curl -XPOST -H "Content-Type: application/json" -d @payload.txt https://wkcfzdti0i.execute-api.us-east-1.amazonaws.com/prod/set
```

get users' blips
```bash
curl -XPOST https://wkcfzdti0i.execute-api.us-east-1.amazonaws.com/prod/get/user/mike@mkdavies.com
curl -XPOST https://wkcfzdti0i.execute-api.us-east-1.amazonaws.com/prod/get/user/rachel@mkdavies.com
```

get a single blip ... update the blipid as needed
```bash
curl -XPOST https://wkcfzdti0i.execute-api.us-east-1.amazonaws.com/prod/get/35f4b75b-3406-4776-a3b8-d6bed6ef938a
```

update a blip
```bash
echo '{"blipid":"35f4b75b-3406-4776-a3b8-d6bed6ef938a","blip":"Blipping update","username":"rachel@mkdavies.com"}' > payload.txt
curl -XPOST -H "Content-Type: application/json" -d @payload.txt https://wkcfzdti0i.execute-api.us-east-1.amazonaws.com/prod/set
```

delete a blip
```bash
echo '{"blipid":"35f4b75b-3406-4776-a3b8-d6bed6ef938a"}' > payload.txt
curl -XPOST -H "Content-Type: application/json" -d @payload.txt https://wkcfzdti0i.execute-api.us-east-1.amazonaws.com/prod/delete
```
