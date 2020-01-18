package queue

import (
	"flag"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

var svc *sqs.SQS
var queueURL *string

func InitQueue() string {
	qName := flag.String("n", "", "Name of the queue used for communication.")
	region := flag.String("r", "us-west-2", "AWS region where the queue exists.")
	flag.Parse()

	if *qName == "" {
		flag.PrintDefaults()
		log.Println("Queue name is required.")
		os.Exit(1)
	}

	sess, _ := session.NewSession(&aws.Config{
		Region: region,
	})

	svc = sqs.New(sess)

	result, err := svc.GetQueueUrl(&sqs.GetQueueUrlInput{
		QueueName: qName,
	})

	queueURL = result.QueueUrl

	if err != nil {
		log.Println("Error: ", err)
		os.Exit(1)
	}

	return *queueURL
}
