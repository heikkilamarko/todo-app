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

	err = c.ScheduleClient().GetHandle(ctx, internal.ScheduleID).Delete(ctx)
	if err != nil {
		log.Fatalln("delete schedule:", err)
	}
}
