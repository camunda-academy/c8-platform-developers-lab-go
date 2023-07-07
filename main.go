package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/camunda/zeebe/clients/go/v8/pkg/entities"
	"github.com/camunda/zeebe/clients/go/v8/pkg/worker"
	"github.com/camunda/zeebe/clients/go/v8/pkg/zbc"
	"log"
	"os"
)

func main() {

	// setup credentials provider
	credsProvider, err := zbc.NewOAuthCredentialsProvider(&zbc.OAuthProviderConfig{
		ClientID:     "XXXXXXXXXXXXXXXX",
		ClientSecret: "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
		Audience:     "zeebe.camunda.io",
	})
	if err != nil {
		panic(err)
	}

	// create client
	client, err := zbc.NewClient(&zbc.ClientConfig{
		GatewayAddress:      "XXXXXXXXXXXXXXXX.XXXX.zeebe.camunda.io:443",
		CredentialsProvider: credsProvider,
	})
	if err != nil {
		panic(err)
	}

	// check connection
	ctx := context.Background()
	response, err := client.NewTopologyCommand().Send(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to " + response.GetBrokers()[1].GetHost())

	// start worker
	jobWorker := client.NewJobWorker().JobType("YOUR_JOB_TYPE").Handler(handleJob).Open()

	// Shut down worker when pushing enter in console
	buf := bufio.NewReader(os.Stdin)
	buf.ReadLine()

	fmt.Println("Shutting down...")

	jobWorker.Close()
	jobWorker.AwaitClose()
}

func handleJob(client worker.JobClient, job entities.Job) {
	jobKey := job.GetKey()

	fmt.Println("Complete job", jobKey, "of type", job.Type)

	ctx := context.Background()
	_, err := client.NewCompleteJobCommand().JobKey(jobKey).Send(ctx)
	if err != nil {
		panic(err)
	}

	log.Println("Successfully completed job")
}