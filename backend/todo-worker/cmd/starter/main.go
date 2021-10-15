package main

import (
	"context"
	"log"
	"todo-worker/workflow"

	"go.temporal.io/sdk/client"
)

func main() {
	c, err := client.NewClient(client.Options{})
	if err != nil {
		log.Fatalln("unable to create temporal client", err)
	}
	defer c.Close()

	options := client.StartWorkflowOptions{
		TaskQueue:    workflow.GreetTaskQueue,
		CronSchedule: "*/1 * * * *",
	}

	_, err = c.ExecuteWorkflow(context.Background(), options, workflow.GreetWorkflow)
	if err != nil {
		log.Fatalln("unable to start cron workflow", err)
	}
}
