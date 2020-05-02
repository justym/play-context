package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()

	parent, cancel := context.WithCancel(ctx)
	go parentFunc(parent, "Canceled Parent")

	cancel()
	time.Sleep(1 * time.Second)
	fmt.Println("FINISH")
}

func parentFunc(ctx context.Context, message string) {
	child, cancel := context.WithCancel(ctx)
	go childFunc(child, "Canceled Child")
	defer cancel()

	for {
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Err(), message)
			return
		}
	}
}

func childFunc(ctx context.Context, message string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Err(), message)
			return
		}
	}
}
