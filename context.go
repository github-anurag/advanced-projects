package main

import (
	"context"
	"fmt"
	"time"
)

func slowOperation(ctx context.Context) {
	for {
		select {
		case <-time.After(1 * time.Second):
			// Simulate doing a piece of work
			fmt.Println("Working...")
		case <-ctx.Done():
			// The context was cancelled or timed out
			fmt.Println("Worker stopping:", ctx.Err())
			return
		}
	}
}

func main() {
	// 1. Create a context that expires after 2.5 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 2500*time.Millisecond)

	// 2. Always defer cancel() to prevent context leaks
	defer cancel()

	// 3. Start the worker
	go slowOperation(ctx)

	// 4. Wait long enough to see the timeout happen
	time.Sleep(4 * time.Second)
	fmt.Println("Main program exiting.")
}
