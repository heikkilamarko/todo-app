package workflow

import (
	"time"

	"go.temporal.io/sdk/workflow"
)

func GreetWorkflow(ctx workflow.Context) error {
	options := workflow.ActivityOptions{
		StartToCloseTimeout: 5 * time.Second,
	}

	ctx = workflow.WithActivityOptions(ctx, options)

	return workflow.ExecuteActivity(ctx, GreetActivity).Get(ctx, nil)
}
