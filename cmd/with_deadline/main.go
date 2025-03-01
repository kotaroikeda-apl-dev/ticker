package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	fmt.Println("Ticker with context deadline started")

	// 今から 5 秒後の時刻を取得
	deadline := time.Now().Add(5 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Deadline reached. Stopping ticker...")
			return
		case t := <-ticker.C:
			fmt.Println("Tick at:", t)
		}
	}
}
