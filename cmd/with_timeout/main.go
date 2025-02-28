package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	fmt.Println("Ticker with context timeout started")

	// `context.WithTimeout` を使い、5秒後に自動でキャンセル
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Timeout reached. Stopping ticker...")
			return
		case t := <-ticker.C:
			fmt.Println("Tick at:", t)
		}
	}
}
