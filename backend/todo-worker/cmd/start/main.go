package main

import (
	"context"
	"log"
	"todo-worker/internal"

	"go.temporal.io/sdk/client"
)

func main() {
	ctx := context.Background()

	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("create temporal client:", err)
	}
	defer c.Close()

	_, err = c.ScheduleClient().Create(ctx, client.ScheduleOptions{
		ID: internal.ScheduleID,
		Spec: client.ScheduleSpec{
			CronExpressions: []string{"*/1 * * * *"},
		},
		Action: &client.ScheduleWorkflowAction{
			ID:        internal.WorkflowID,
			Workflow:  internal.RemoveTodosWorkflow,
			TaskQueue: internal.TaskQueue,
		},
	})
	if err != nil {
		log.Fatalln("create schedule:", err)
	}
}
