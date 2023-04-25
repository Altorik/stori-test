package handler

import (
	"context"
	"log"
	"process-file/internal"
	"process-file/internal/process"
	"process-file/internal/queue"
)

func ProcessHandler(ctx context.Context, fileService *process.FileService, queueService *queue.QueueService) error {
	log.Printf("test")
	var fileToProcess internal.File
	err := fileToProcess.NewFileToProcess("", fileService.FileName)
	if err != nil {
		return err
	}
	summary, err := fileService.GetSummary(fileToProcess)
	if err != nil {
		return err
	}
	err = queueService.Send(summary)
	if err != nil {
		return err
	}
	return nil
}
