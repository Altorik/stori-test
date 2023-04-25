package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"log"
	"send-email/cmd/sqs/bootstrap"
)

func handler(ctx context.Context, sqsEvent events.SQSEvent) {
	if err := bootstrap.Run(ctx, sqsEvent); err != nil {
		log.Panic(err)
	}
}
func main() {
	lambda.Start(handler)
	//err := handlery(events.APIGatewayProxyRequest{})
	//if err != nil {
	//	return
	//}
}
