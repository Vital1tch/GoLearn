package concurency

import (
	"context"
	"fmt"
)

func processRequest(ctx context.Context) {
	if userID, ok := ctx.Value("UserID").(string); ok {
		fmt.Printf("UserID: %s", userID)
	} else {
		fmt.Println("User not found")
	}
}
