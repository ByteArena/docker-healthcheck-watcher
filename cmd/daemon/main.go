package main

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	sdk "github.com/aws/aws-sdk-go/service/sns"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"

	"github.com/bytearena/docker-healthcheck-watcher/template"
	t "github.com/bytearena/docker-healthcheck-watcher/types"
)

func newMessage(str string) *string {
	return &str
}

func newTopicArn(str string) *string {
	return &str
}

func onContainerHealthCheckFailure(sns *sdk.SNS, service string) {
	errorMessage := t.ErrorMessage{
		ServiceName:   service,
		ServiceStatus: "unhealthy (running)",
		Log:           "",
	}

	message := template.MakeTemplate(errorMessage)

	topicArn := os.Getenv("TOPIC_ARN")
	msg := sdk.PublishInput{
		TopicArn: newTopicArn(topicArn),
		Message:  newMessage(message),
	}

	output, err := sns.Publish(&msg)

	if err != nil {
		log.Panicln(err)
	}

	log.Println(service, "failure, sent message", output.MessageId)
}

func NewSNSClient() *sdk.SNS {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			Credentials: credentials.NewEnvCredentials(),
			Region:      aws.String("eu-west-1"),
		},
	}))

	return sdk.New(sess)
}

func main() {
	cli, clientError := client.NewEnvClient()
	ctx := context.Background()

	if clientError != nil {
		log.Panicln(clientError)
	}

	stream, err := cli.Events(ctx, types.EventsOptions{})

	snsClient := NewSNSClient()

	for {
		select {
		case msg := <-err:
			log.Panicln(msg)
		case msg := <-stream:
			if msg.Action == "health_status: unhealthy" {
				onContainerHealthCheckFailure(snsClient, msg.Actor.Attributes["image"])
			}
		}
	}
}
