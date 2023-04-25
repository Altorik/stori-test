package main

import (
	"baz/pacl-task-service-orquestador/cmd/sqs"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(sqs.Start)
}
