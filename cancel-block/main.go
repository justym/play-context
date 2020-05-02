package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(10*time.Second))
	defer cancel()

	go sayBye(ctx)

	select {
	case <-ctx.Done():
		fmt.Println(ctx.Err())
		break
	}
}

func sayBye(ctx context.Context) {
	for {
		fmt.Println(ctx.Deadline())
		time.Sleep(1 * time.Second)
	}
}
