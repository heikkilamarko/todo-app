package main

import (
	"context"
	"log"
	"todo-worker/internal"

	"go.temporal.io/sdk/client"
)

func main() {
	c, err := client.NewClient(client.Options{})
	if err != nil {
		log.Fatalln("unable to create temporal client", err)
	}
	defer c.Close()

	options := client.StartWorkflowOptions{
		ID:           internal.WorkflowID,
		TaskQueue:    internal.TaskQueueWorker,
		CronSchedule: "*/1 * * * *",
	}

	_, err = c.ExecuteWorkflow(context.Background(), options, internal.RemoveTodosWorkflow)
	if err != nil {
		log.Fatalln("unable to start cron workflow", err)
	}
}
