package main

import (
	"context"
	"log"

	"email-reminder/workflows"

	"go.temporal.io/sdk/client"
)

// ! FOR TEST WITHOUT API
// ! RUN THIS FILE FOR TEST
func main() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to connect to Temporal:", err)
	}
	defer c.Close()

	workflowOptions := client.StartWorkflowOptions{
		ID:        "email_reminder_workflow_001",
		TaskQueue: "EMAIL_TASK_QUEUE",
	}

	email := "user@example.com"
	delay := 10 // seconds

	log.Println("🚀 Starting workflow to remind:", email)
	we, err := c.ExecuteWorkflow(context.Background(), workflowOptions, workflows.EmailReminderWorkflow, email, delay)
	if err != nil {
		log.Fatalln("Unable to execute workflow:", err)
	}

	log.Println("✅ Workflow started", "WorkflowID:", we.GetID(), "RunID:", we.GetRunID())
}
