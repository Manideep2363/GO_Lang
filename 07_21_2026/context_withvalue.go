package main

import (
	"context"
	"fmt"
)

type contextKey string

const userIDKey contextKey = "userID"

func process(ctx context.Context) {
	userId := ctx.Value(userIKey)
	fmt.Println(userId)

}

func main() {

	parentCtx := context.Background()
	ctx := context.WithValue(parentCtx, userIDKey, 101)

	process(ctx)
}
