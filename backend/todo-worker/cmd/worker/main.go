package main

import (
	"log"
	"todo-worker/workflow"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {
	c, err := client.NewClient(client.Options{})
	if err != nil {
		log.Fatalln("unable to create temporal client", err)
	}
	defer c.Close()

	w := worker.New(c, workflow.GreetTaskQueue, worker.Options{})
	w.RegisterWorkflow(workflow.GreetWorkflow)
	w.RegisterActivity(workflow.GreetActivity)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("unable to start worker", err)
	}
}
