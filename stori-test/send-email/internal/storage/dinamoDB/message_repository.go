package dinamoDB

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	email "send-email/internal"
)

type ProcessedMessageRepositorio struct {
	db        *dynamodb.Client
	tableName string
}

func NuevoProcessedMessageRepositorio(db *dynamodb.Client, tableName string) *ProcessedMessageRepositorio {
	return &ProcessedMessageRepositorio{
		db:        db,
		tableName: tableName,
	}
}

func (r *ProcessedMessageRepositorio) Guardar(message email.ProcessMessage) error {
	nuevoMensage, err := attributevalue.MarshalMap(noSqlMessage{
		MessageID: message.MessageId,
		DateEvent: message.DateEvent,
	})
	if err != nil {
		return err
	}

	input := &dynamodb.PutItemInput{
		Item:      nuevoMensage,
		TableName: aws.String(r.tableName),
	}
	_, err = r.db.PutItem(context.TODO(), input)
	if err != nil {
		return err
	}

	return nil
}

func (r *ProcessedMessageRepositorio) Buscar(message email.ProcessMessage) error {
	llaveBusqueda, err := attributevalue.MarshalMap(noSqlMessage{
		MessageID: message.MessageId,
		DateEvent: message.DateEvent,
	})
	if err != nil {
		return err
	}
	data, err := r.db.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: aws.String(r.tableName),
		Key:       llaveBusqueda,
	})

	if err != nil {
		return err
	}
	if data.Item != nil {
		return fmt.Errorf("repetido")
	}
	return nil
}
