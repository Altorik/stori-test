package bootstrap

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/kelseyhightower/envconfig"
	"log"
	mySqs "process-file/internal/bus/Sqs"
	"process-file/internal/platform/lambda"
	"process-file/internal/process"
	"process-file/internal/queue"
	myS3 "process-file/internal/storage/s3"
	"time"
)

func Run() error {
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
	sqsUrl := fmt.Sprintf("https://sqs.%s.amazonaws.com/%s/%s", configEnv.Region, configEnv.CuentaAws, configEnv.QueueName)

	log.Println("CONFIG")
	log.Println(configEnv)
	log.Println(sqsUrl)
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(configEnv.Region))
	if err != nil {
		return err
	}
	s3Client := s3.NewFromConfig(cfg)
	sqsClient := sqs.NewFromConfig(cfg)

	s3Repository := myS3.NewS3Repository(s3Client, configEnv.BucketName)
	fileService := process.NewFileService(s3Repository, configEnv.FileName)
	sqsRepository := mySqs.NewSqsRepository(sqsClient, sqsUrl)
	queueService := queue.NewQueueService(sqsRepository)

	log.Println("Starting Stori Service")
	ctx, srv := lambda.New(context.Background(), fileService, queueService)
	return srv.Run(ctx)
}

type configMicro struct {
	TimeZone   string `default:"America/Mexico_City"`
	CuentaAws  string `default:"679840262051"`
	Region     string `default:"us-east-1"`
	BucketName string `default:"test-stori"`
	QueueName  string `default:"sqs-process-files"`
	FileName   string `default:"csv_client_2.csv"`
}
