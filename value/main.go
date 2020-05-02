package main

import (
	"context"
	"fmt"
	"time"
)

const (
	keyA   = "a"
	keyB   = "b"
	keyC   = "c"
	twosec = 2 * time.Second
	red    = "\033[1;31m%s\033[0m"
	blue   = "\033[1;34m%s\033[0m"
	yellow = "\033[1;33m%s\033[0m"
)

func main() {
	ctx := context.Background()

	ctxWithA := context.WithValue(ctx, keyA, 1)
	ctxWithB := context.WithValue(ctxWithA, keyB, 2)
	ctxWithC := context.WithValue(ctxWithB, keyC, 3)
	ctxOverC := context.WithValue(ctxWithC, keyC, 10)

	go agentA(ctxWithA)
	go agentB(ctxWithB)
	go agentC(ctxWithC)
	go agentC(ctxOverC)

	time.Sleep(6 * time.Second)
}

func agentA(ctx context.Context) {
	for {
		fmt.Printf(red, "== This is agentA ==\n")
		fmt.Printf("%s: %d\n", keyA, ctx.Value(keyA))
		time.Sleep(twosec)
	}
}

func agentB(ctx context.Context) {
	for {
		fmt.Printf(blue, "== This is agentB ==\n")
		fmt.Printf(
			"%s: %d, %s: %d\n",
			keyA, ctx.Value(keyA),
			keyB, ctx.Value(keyB),
		)
		time.Sleep(twosec)
	}
}

func agentC(ctx context.Context) {
	for {
		fmt.Printf(yellow, "== This is agentC ==\n")
		fmt.Printf(
			"%s: %d, %s: %d, %s: %d\n",
			keyA, ctx.Value(keyA),
			keyB, ctx.Value(keyB),
			keyC, ctx.Value(keyC),
		)
		time.Sleep(twosec)
	}
}
