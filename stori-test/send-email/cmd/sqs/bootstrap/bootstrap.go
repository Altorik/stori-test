package bootstrap

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/sesv2"
	"log"
	emailSer "send-email/internal/email"
	"send-email/internal/inserting"
	"send-email/internal/platform/server"
	"send-email/internal/platform/server/email"
	"send-email/internal/storage/dinamoDB"

	"github.com/kelseyhightower/envconfig"
	"time"
)

func Run(_ context.Context, sqsEvent events.SQSEvent) error {
	var configEnv configMicro
	err := envconfig.Process("STORI", &configEnv)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	if tz := configEnv.TimeZone; tz != "" {
		time.Local, err = time.LoadLocation(tz)
		if err != nil {
			return err
		}
	}

	log.Println("CONFIG")
	log.Println(configEnv)
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(configEnv.Region))
	if err != nil {
		panic(err)
	}
	dynamoDBConfiguracion := dynamodb.NewFromConfig(cfg)
	sesv2Config := sesv2.NewFromConfig(cfg)

	processedMessageRepositorio := dinamoDB.NuevoProcessedMessageRepositorio(dynamoDBConfiguracion, configEnv.TableName)
	processedMessageService := inserting.NuevoProcessedMessageService(processedMessageRepositorio)
	emailRepository := email.NewEmailRepository(sesv2Config, configEnv.EmailName, configEnv.TemplateName)
	emailService := emailSer.NewEmailService(emailRepository)
	srv := server.New(sqsEvent, processedMessageService, emailService)
	return srv.Run()
}

type configMicro struct {
	TimeZone     string `default:"America/Mexico_City"`
	CuentaAws    string `default:"679840262051"`
	Region       string `default:"us-east-1"`
	QueueName    string `default:"sqs-process-files"`
	TemplateName string `default:"test-stori"`
	EmailName    string `default:"test-stori"`
	TableName    string `default:"test-stori"`
}
