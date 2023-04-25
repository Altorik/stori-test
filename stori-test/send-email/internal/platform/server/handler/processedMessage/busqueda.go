package processedMessage

import (
	"github.com/aws/aws-lambda-go/events"
	message "send-email/internal"
	"send-email/internal/inserting"
)

func Search(mensaje events.SQSMessage, serviceDynamo inserting.ProcessMessageService) error {
	err := serviceDynamo.ProcessMessageRepositorio.Buscar(message.ProcessMessage{
		MessageId: mensaje.MessageId,
		DateEvent: mensaje.Md5OfBody,
	})
	if err != nil {
		return err
	}
	return nil
}
