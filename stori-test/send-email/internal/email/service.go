package email

import (
	"math"
	"send-email/internal"
	"send-email/internal/platform/server/email"
)

type EmailService struct {
	EmailRepository *email.EmailRepository
}

func NewEmailService(EmailRepository *email.EmailRepository) *EmailService {
	return &EmailService{
		EmailRepository: EmailRepository,
	}
}

func (e *EmailService) Send(summary internal.Summary) error {
	summaryTemplate := convertToSummaryTemplate(summary)
	err := e.EmailRepository.Send(summaryTemplate)
	if err != nil {
		return err
	}
	return nil
}

func convertToSummaryTemplate(summary internal.Summary) (result internal.SummaryTemplate) {
	transactionsByMonth := make([]internal.TransactionByMonth, 0)
	for month, txns := range summary.TransactionsByMonth {
		if txns < 1 {
			continue
		}
		transactionsByMonth = append(transactionsByMonth, internal.TransactionByMonth{
			NameMonth:    internal.Months[month],
			Transactions: txns,
		})
	}
	result.TotalBalance = roundNearest(summary.TotalBalance)
	result.AverageDebitAmount = roundNearest(summary.AverageDebitAmount)
	result.AverageCreditAmount = roundNearest(summary.AverageCreditAmount)
	result.TransactionsByMonth = transactionsByMonth
	return result
}

func roundNearest(value float64) float64 {
	return math.Round(value*100) / 100
}
