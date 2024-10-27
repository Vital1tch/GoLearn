package concurency

import (
	"context"
	"fmt"
)

func processToken(ctx context.Context, key string, authToken string) {
	if userID, ok := ctx.Value(key).(string); ok {
		fmt.Printf("UserId: <%s>\n", userID)
	} else {
		fmt.Println("Ключ не найден")
	}

	if token, ok := ctx.Value(authToken).(string); ok {
		fmt.Printf("Token: <%s>\n", token)
	} else {
		fmt.Println("Токен не найден")
	}
}
