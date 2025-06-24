package main

import (
	"log"

	"email-reminder/activities"
	"email-reminder/workflows"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to connect to Temporal:", err)
	}
	defer c.Close()

	w := worker.New(c, "EMAIL_TASK_QUEUE", worker.Options{})

	w.RegisterWorkflow(workflows.EmailReminderWorkflow)
	w.RegisterActivity(activities.SendEmail)

	log.Println("👷 Worker started.")
	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Worker failed:", err)
	}
}
