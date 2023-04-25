package lambda

import (
	"context"
	"log"
	"process-file/internal/platform/lambda/handler"
	"process-file/internal/process"
	"process-file/internal/queue"
)

type Lambda struct {
	fileService  *process.FileService
	queueService *queue.QueueService
}

func New(ctx context.Context, fileService *process.FileService, queueService *queue.QueueService) (context.Context, Lambda) {
	srv := Lambda{
		fileService:  fileService,
		queueService: queueService,
	}

	return ctx, srv
}

func (s *Lambda) Run(ctx context.Context) error {
	err := handler.ProcessHandler(ctx, s.fileService, s.queueService)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}
