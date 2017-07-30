package sqscheck

import (
	"fmt"
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

// Run function prepare and run list of checks by provided config arg
func Run(config *Config) {
	sqsClient := getSQSClient(config.AwsAccessID, config.AwsAccessSecret, config.AwsRegion)

	for _, queue := range config.Queues {
		check := config.DefaultCheckSpec
		for _, customCheck := range config.CustomChecks {
			if contain(customCheck.Queues, queue) {
				check = customCheck
			}
		}
		// Run check for a queue
		runCheck(queue, check, config, sqsClient)

		// Run check for the respective error queue
		runCheck(queue+"_errors", CheckSpec{
			CriticalThreashold: 1,
		}, config, sqsClient)
	}
}

func queueLength(queueName string, region, account string, sqsClient *sqs.SQS) (int, error) {
	queueSize := -1
	queueURL := fmt.Sprintf("https://sqs.%s.amazonaws.com/%s/%s", region, account, queueName)

	resp, err := sqsClient.GetQueueAttributes(
		&sqs.GetQueueAttributesInput{
			QueueUrl: aws.String(queueURL),
			AttributeNames: []*string{
				aws.String("ApproximateNumberOfMessages"),
			},
		})

	for attrib := range resp.Attributes {
		prop := resp.Attributes[attrib]
		i, _ := strconv.Atoi(*prop)
		queueSize = i
	}

	return queueSize, err

	// queueSize = rand.Intn(30)
	// return queueSize, err
}

func contain(array []string, match string) bool {
	for _, i := range array {
		if i == match {
			return true
		}
	}
	return false
}

func runCheck(queue string, check CheckSpec, config *Config, sqs *sqs.SQS) {
	length, err := queueLength(queue, config.AwsRegion, config.AwsAccountNum, sqs)
	if err != nil {
		fmt.Println("runDefaultCheck error has occured")
		fmt.Println(err)
	}
	// fmt.Printf("QUEUE: %s, LEN: %d, CRIT: %d, WARN: %d \n", queue, length, check.CriticalThreashold, check.WarningThreashold)
	if length > check.CriticalThreashold {
		fmt.Printf("[CRITICAL] %s queue contain %d messages\n", queue, length)
		return
	}
	if length > check.WarningThreashold {
		fmt.Printf("[WARNING] %s queue contain %d messages\n", queue, length)
		return
	}
}

func getSQSClient(accessID string, accessSecret string, region string) *sqs.SQS {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewStaticCredentials(accessID, accessSecret, ""),
	})
	if err != nil {
		fmt.Println(err)
	}

	// Create a SQS service client.
	sqsClient := sqs.New(sess)
	return sqsClient
}
