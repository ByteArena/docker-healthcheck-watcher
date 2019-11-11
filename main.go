package main

import (
	"context"
	"log"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"

	slack "github.com/bytearena/docker-healthcheck-watcher/integration"
	"github.com/bytearena/docker-healthcheck-watcher/template"
	t "github.com/bytearena/docker-healthcheck-watcher/types"
	"github.com/stratsys/go-common/env"
)

func onContainerDieFailure(service string, exitCode string) {
	errorMessage := t.ErrorMessage{
		Emoji:         ":red_circle:",
		ServiceName:   service,
		ServiceStatus: "died (exited with code " + exitCode + ")",
		Log:           "",
	}

	message := template.MakeTemplate(errorMessage)

	output := slack.Publish(message)

	log.Println(service, "failure, sent message", output)
}

func onContainerHealthy(service string) {
	errorMessage := t.ErrorMessage{
		Emoji:         ":+1:",
		ServiceName:   service,
		ServiceStatus: "ok",
		Log:           "",
	}

	message := template.MakeTemplate(errorMessage)

	output := slack.Publish(message)

	log.Println(service, "sent message", output)
}

func onContainerHealthCheckFailure(service string) {
	errorMessage := t.ErrorMessage{
		Emoji:         "🚨",
		ServiceName:   service,
		ServiceStatus: "unhealthy (running)",
		Log:           "",
	}

	message := template.MakeTemplate(errorMessage)

	output := slack.Publish(message)

	log.Println(service, "failure, sent message", output)
}

func main() {
	kvps := make(map[string]string)
	for _, file := range os.Args[1:] {
		if err := env.ReadFiles(file, kvps); err != nil {
			panic(err)
		}
	}

	for k, v := range kvps {
		os.Setenv(k, v)
	}

	cli, clientError := client.NewEnvClient()
	ctx := context.Background()

	if clientError != nil {
		log.Panicln(clientError)
	}

	stream, err := cli.Events(ctx, types.EventsOptions{})

	for {
		select {
		case msg := <-err:
			log.Panicln(msg)
		case msg := <-stream:
			if msg.Action == "health_status: unhealthy" {
				onContainerHealthCheckFailure(msg.Actor.Attributes["image"])
			}

			if msg.Action == "health_status: healthy" {
				onContainerHealthy(msg.Actor.Attributes["image"])
			}

			if msg.Action == "die" {
				exitCode := msg.Actor.Attributes["exitCode"]

				if exitCode != "0" {
					onContainerDieFailure(msg.Actor.Attributes["image"], exitCode)
				}
			}
		}
	}
}
