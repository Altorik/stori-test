package queue

import (
	client "process-file/internal"
	mySqs "process-file/internal/bus/Sqs"
)

type QueueService struct {
	sqsRepository *mySqs.Repository
}

func NewQueueService(sqsRepository *mySqs.Repository) *QueueService {
	return &QueueService{
		sqsRepository: sqsRepository,
	}
}

func (f *QueueService) Send(info client.Summary) error {
	err := f.sqsRepository.Send(info)
	if err != nil {
		return err
	}
	return nil
}
