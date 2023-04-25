package dinamoDB

type noSqlMessage struct {
	MessageID string `dynamodbav:"MessageId"`
	DateEvent string `dynamodbav:"DateEvent"`
}
