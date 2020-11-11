# GO-KIT EXAMPLE

API to ilustrate a basic usage for go-kit. Based in the tutorial from https://github.com/tensor-programming/go-kit-tutorial

## RUN

First run localstack for dynamoDB service

``` bash
docker-compose up
```

If the table does not exist, just go ahead and create it

``` bash
aws --endpoint-url=http://localhost:4569 dynamodb create-table \
    --table-name users \
    --attribute-definitions AttributeName=UserId,AttributeType=S \
    --key-schema AttributeName=UserId,KeyType=HASH \
    --provisioned-throughput ReadCapacityUnits=50,WriteCapacityUnits=50 \
    --region us-east-1
```

Then just

``` bash
go run main.go
```
