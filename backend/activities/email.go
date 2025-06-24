package activities

import (
	"context"
	"fmt"
)

func SendEmail(ctx context.Context, email string) error {
	//print this msg after delay
	fmt.Println("📧 Sending email to:", email)
	return nil
}
