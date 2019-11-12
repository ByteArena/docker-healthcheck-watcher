package main

import (
	"context"
	"log"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/events"
	"github.com/docker/docker/client"

	"github.com/bytearena/docker-healthcheck-watcher/alertnotification"
	"github.com/stratsys/go-common/env"
)

func onContainerDieFailure(service string, exitCode string, attributes map[string]string) {
	if err := alertnotification.NewMsTeam("ff5864", service, "died (exited with code "+exitCode+")", attributes).Send(); err != nil {
		log.Panicln(err)
	}
}

func onContainerHealthy(service string, attributes map[string]string) {
	if err := alertnotification.NewMsTeam("90ee90", service, "ok", attributes).Send(); err != nil {
		log.Panicln(err)
	}
}

func onContainerHealthCheckFailure(service string, attributes map[string]string) {
	if err := alertnotification.NewMsTeam("ff5864", service, "unhealthy (running)", attributes).Send(); err != nil {
		log.Panicln(err)
	}
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
			go func(msg events.Message) {
				handleMessage(msg)
			}(msg)
		}
	}
}

func handleMessage(msg events.Message) {
	if msg.Action == "health_status: unhealthy" {
		onContainerHealthCheckFailure(getServiceName(msg.Actor.Attributes), msg.Actor.Attributes)
	}

	if msg.Action == "health_status: healthy" {
		onContainerHealthy(getServiceName(msg.Actor.Attributes), msg.Actor.Attributes)
	}

	if msg.Action == "die" {
		exitCode := msg.Actor.Attributes["exitCode"]

		if exitCode != "0" {
			onContainerDieFailure(getServiceName(msg.Actor.Attributes), exitCode, msg.Actor.Attributes)
		}
	}
}

func getServiceName(attributes map[string]string) string {
	if nm, ok := attributes["com.docker.swarm.service.name"]; ok {
		return nm
	}

	return attributes["image"]
}
