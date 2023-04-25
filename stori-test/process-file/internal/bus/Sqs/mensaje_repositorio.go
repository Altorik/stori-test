package Sqs

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"log"
	email "process-file/internal"
)

type Repository struct {
	sqs      *sqs.Client
	urlQueue string
}

func NewSqsRepository(sqs *sqs.Client, urlQueue string) *Repository {
	return &Repository{
		sqs:      sqs,
		urlQueue: urlQueue,
	}
}

func (o *Repository) Send(mensaje email.Summary) error {
	mensajeModoTexto, err := json.Marshal(mensaje)
	_, err = o.sqs.SendMessage(context.Background(), &sqs.SendMessageInput{
		MessageBody: aws.String(string(mensajeModoTexto)),
		QueueUrl:    aws.String(o.urlQueue),
	})
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}
