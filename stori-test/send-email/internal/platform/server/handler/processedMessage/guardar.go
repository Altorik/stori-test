package processedMessage

import (
	"github.com/aws/aws-lambda-go/events"
	message "send-email/internal"
	"send-email/internal/inserting"
)

func Save(mensaje events.SQSMessage, serviceDynamo inserting.ProcessMessageService) error {
	err := serviceDynamo.ProcessMessageRepositorio.Guardar(message.ProcessMessage{
		MessageId: mensaje.MessageId,
		DateEvent: mensaje.Md5OfBody,
	})
	if err != nil {
		return err
	}
	return nil
}
