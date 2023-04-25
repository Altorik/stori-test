package email

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sesv2"
	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"
	email "send-email/internal"
)

type EmailRepository struct {
	sesClient    *sesv2.Client
	emailToUse   []string
	templateName string
}

func NewEmailRepository(sesClient *sesv2.Client, emailToUse, templateName string) *EmailRepository {
	a2 := make([]string, 1)
	a2[0] = emailToUse
	return &EmailRepository{
		sesClient:    sesClient,
		emailToUse:   a2,
		templateName: templateName,
	}
}

func (r *EmailRepository) Send(summary email.SummaryTemplate) error {
	summaryText, err := json.Marshal(summary)
	if err != nil {
		return err
	}
	emailToSend := &types.EmailContent{
		Template: &types.Template{
			//TemplateArn:  nil,
			TemplateData: aws.String(string(summaryText)),
			TemplateName: aws.String(r.templateName),
		},
	}
	input := &sesv2.SendEmailInput{
		Content: emailToSend,
		Destination: &types.Destination{
			ToAddresses: r.emailToUse,
		},
		FromEmailAddress: aws.String(r.emailToUse[0]),
	}
	_, err = r.sesClient.SendEmail(context.TODO(), input)
	if err != nil {
		return err
	}

	return nil
}
