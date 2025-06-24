package workflows

import (
	"time"

	"email-reminder/activities"

	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

func EmailReminderWorkflow(ctx workflow.Context, email string, delaySec int) error {
	logger := workflow.GetLogger(ctx)

	ao := workflow.ActivityOptions{
		StartToCloseTimeout: time.Minute,
		RetryPolicy: &temporal.RetryPolicy{
			InitialInterval:    time.Second * 5,
			MaximumAttempts:    3,
			BackoffCoefficient: 2.0,
		},
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	logger.Info("⏳ Waiting before sending email...")
	_ = workflow.Sleep(ctx, time.Duration(delaySec)*time.Second)

	err := workflow.ExecuteActivity(ctx, activities.SendEmail, email).Get(ctx, nil)
	if err != nil {
		logger.Error("❌ Failed to send email", "error", err)
	}
	return err
}
