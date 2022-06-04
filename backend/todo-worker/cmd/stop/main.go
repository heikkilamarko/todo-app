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

	err = c.CancelWorkflow(context.Background(), internal.WorkflowID, "")
	if err != nil {
		log.Fatalln("unable to cancel workflow execution", err)
	}

	log.Printf("workflow execution cancelled (ID: '%s')", internal.WorkflowID)
}
