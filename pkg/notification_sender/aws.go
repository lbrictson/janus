package notification_sender

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/eventbridge"
	"github.com/aws/aws-sdk-go-v2/service/eventbridge/types"
	"github.com/aws/aws-sdk-go-v2/service/sns"
)

type AWSEvent struct {
	NewNotificationInput
	EventSourceName string `json:"event_source_name"`
}

func SendEventBridgeNotification(input NewNotificationInput, AWSAccessKeyID string, AWSSecretAccessKey string, region string, eventBusName string, eventSource string, eventDetail string) error {
	creds := credentials.NewStaticCredentialsProvider(AWSAccessKeyID, AWSSecretAccessKey, "")

	// Load AWS configuration
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithCredentialsProvider(creds),
		config.WithRegion(region),
	)
	if err != nil {
		return err
	}

	a := AWSEvent{
		NewNotificationInput: input,
		EventSourceName:      "janus",
	}

	b, err := json.Marshal(a)
	if err != nil {
		return err
	}
	s := string(b)
	client := eventbridge.NewFromConfig(cfg)

	i := &eventbridge.PutEventsInput{
		Entries: []types.PutEventsRequestEntry{
			{
				Detail:       &s,
				DetailType:   &eventDetail,
				EventBusName: &eventBusName,
				Source:       &eventSource,
			},
		},
	}

	_, err = client.PutEvents(context.TODO(), i)
	return err
}

func SendSNSNotification(input NewNotificationInput, AWSAccessKeyID string, AWSSecretAccessKey string, region string, topicARN string) error {
	creds := credentials.NewStaticCredentialsProvider(AWSAccessKeyID, AWSSecretAccessKey, "")

	// Load AWS configuration
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithCredentialsProvider(creds),
		config.WithRegion(region),
	)
	if err != nil {
		return err
	}

	a := AWSEvent{
		NewNotificationInput: input,
		EventSourceName:      "janus",
	}

	// Create SNS client
	client := sns.NewFromConfig(cfg)

	// Marshal message to JSON
	messageBytes, err := json.Marshal(a)
	if err != nil {
		return err
	}

	m := string(messageBytes)

	// Publish message
	_, err = client.Publish(context.TODO(), &sns.PublishInput{
		Message:  &m,
		TopicArn: &topicARN,
	})

	return err
}
