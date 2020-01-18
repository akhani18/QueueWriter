package queue

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func Send(payload []byte) (string, error) {
	result, err := svc.SendMessage(&sqs.SendMessageInput{
		MessageBody: aws.String(string(payload)),
		QueueUrl:    queueURL,
	})

	if err != nil {
		log.Println("Error: ", err)
		return "", err
	}

	return *result.MessageId, nil
}
