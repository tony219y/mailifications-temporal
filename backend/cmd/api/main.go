package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"email-reminder/workflows"

	"go.temporal.io/sdk/client"
)

type ReminderRequest struct {
	Email string `json:"email"`
	Delay int    `json:"delay"` // วินาที
}

func main() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalf("Temporal client connect failed: %v", err)
	}
	defer c.Close()

	reminderHandler := func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
			return
		}

		var req ReminderRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil || req.Email == "" || req.Delay < 0 {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}

		workflowOptions := client.StartWorkflowOptions{
			ID:        "reminder_" + req.Email + time.Now().Format("20060102150405"),
			TaskQueue: "EMAIL_TASK_QUEUE",
		}

		we, err := c.ExecuteWorkflow(context.Background(), workflowOptions, workflows.EmailReminderWorkflow, req.Email, req.Delay)
		if err != nil {
			http.Error(w, "Failed to start workflow: "+err.Error(), http.StatusInternalServerError)
			return
		}

		resp := map[string]string{
			"workflow_id": we.GetID(),
			"run_id":      we.GetRunID(),
			"status":      "started",
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}

	// CORS middleware
	http.Handle("/reminder", corsMiddleware(http.HandlerFunc(reminderHandler)))

	log.Println("Backend API running on :8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}

// CORS Function
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*") // หรือระบุ origin ก็ได้
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
