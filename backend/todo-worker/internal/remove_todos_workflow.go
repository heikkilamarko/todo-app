package internal

import (
	"time"

	"go.temporal.io/sdk/workflow"
)

func RemoveTodosWorkflow(ctx workflow.Context) error {
	options := workflow.ActivityOptions{
		StartToCloseTimeout: time.Minute,
	}

	ctx = workflow.WithActivityOptions(ctx, options)

	a := &Activities{}

	return workflow.ExecuteActivity(ctx, a.RemoveTodos).Get(ctx, nil)
}
