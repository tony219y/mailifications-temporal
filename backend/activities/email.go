package activities

import (
	"context"
	"fmt"
)

func SendEmail(ctx context.Context, email string) error {
	fmt.Println("📧 Sending email to:", email)
	// Mock delay or actual email sending here
	return nil
}
