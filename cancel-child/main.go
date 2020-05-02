package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	parent, cancel := context.WithCancel(ctx)
	go func() {
		if err := childFunc(parent); err != nil {
			fmt.Printf("!! ERROR in parent !!\n")
			cancel()
			return
		}
	}()

	for {
		select {
		case <-parent.Done():
			fmt.Println(parent.Err(), "Parent is done")
			return
		default:
			fmt.Printf("Parent process is working\n")
		}
		time.Sleep(1 * time.Second)
	}
}

func childFunc(ctx context.Context) error {
	child, cancel := context.WithCancel(ctx)
	go func() {
		if err := throwError(); err != nil {
			fmt.Printf("!! ERROR in child !!\n")
			cancel()
		}
	}()

	for {
		select {
		case <-child.Done():
			fmt.Println(child.Err(), "Child is done")
			return fmt.Errorf("")
		default:
			fmt.Printf("Child process is working\n")
		}
		time.Sleep(1 * time.Second)
	}
}

func throwError() error {
	time.Sleep(3 * time.Second)
	return fmt.Errorf("Got Error")
}
