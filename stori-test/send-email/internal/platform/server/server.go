package server

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"log"
	"send-email/internal"
	"send-email/internal/email"
	"send-email/internal/inserting"
	"send-email/internal/platform/server/handler/newEmail"
	"send-email/internal/platform/server/handler/processedMessage"
)

type Server struct {
	mensajes        events.SQSEvent
	dynamoDBService inserting.ProcessMessageService
	emailService    *email.EmailService
}

func New(mensajes events.SQSEvent, processedMessageService inserting.ProcessMessageService, emailService *email.EmailService) Server {
	srv := Server{
		mensajes:        mensajes,
		dynamoDBService: processedMessageService,
		emailService:    emailService,
	}

	return srv
}

func (s *Server) Run() error {
	log.Println("Start Service")

	for _, actualMessage := range s.mensajes.Records {
		var newMessage internal.Summary
		err := json.Unmarshal([]byte(actualMessage.Body), &newMessage)
		if err != nil {
			return err
		}
		err = processedMessage.Search(actualMessage, s.dynamoDBService)
		if err != nil {
			log.Println(err)
			return err
		}
		err = newEmail.SendHandler(newMessage, s.emailService)
		if err != nil {
			log.Println(err)
			return err
		}

		err = processedMessage.Save(actualMessage, s.dynamoDBService)
		if err != nil {
			log.Println(err)
			return err
		}
	}
	return nil
}
