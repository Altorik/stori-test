package inserting

import (
	queue "send-email/internal"
)

type ProcessMessageService struct {
	ProcessMessageRepositorio queue.ProcessMessageRepositorio
}

func NuevoProcessedMessageService(processMessageRepository queue.ProcessMessageRepositorio) ProcessMessageService {
	return ProcessMessageService{
		ProcessMessageRepositorio: processMessageRepository,
	}
}
