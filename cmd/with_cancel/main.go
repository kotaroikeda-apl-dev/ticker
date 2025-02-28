package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	fmt.Println("Ticker with context cancellation started")

	// `context.WithCancel` を作成
	ctx, cancel := context.WithCancel(context.Background())

	// 1秒ごとに tick を送る Ticker
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	// 3秒後にキャンセルを実行
	go func() {
		time.Sleep(3 * time.Second)
		fmt.Println("Cancelling context...")
		cancel() // `context` をキャンセル
	}()

	for {
		select {
		case <-ctx.Done(): // `context` のキャンセルを受信
			fmt.Println("Context cancelled. Stopping ticker...")
			return
		case t := <-ticker.C:
			fmt.Println("Tick at:", t)
		}
	}
}
