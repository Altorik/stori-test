package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"log"
	bootstrap "process-file/cmd/lambda/boostrap"

	"github.com/aws/aws-lambda-go/events"
)

func handler(request events.APIGatewayProxyRequest) error {
	err := bootstrap.Run()
	if err != nil {
		log.Printf(err.Error())
		return err
	}
	return nil
}

func main() {
	lambda.Start(handler)
	//err := handlery(events.APIGatewayProxyRequest{})
	//if err != nil {
	//	return
	//}
}
