package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	fmt.Println("Multiple Tickers with Context")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// 2つの Ticker を作成
	ticker1 := time.NewTicker(1 * time.Second)
	ticker2 := time.NewTicker(2 * time.Second)
	defer ticker1.Stop()
	defer ticker2.Stop()

	// 5秒後にキャンセル
	go func() {
		time.Sleep(5 * time.Second)
		fmt.Println("Cancelling context...")
		cancel()
	}()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Context cancelled. Stopping all tickers...")
			return
		case t := <-ticker1.C:
			fmt.Println("Ticker1 at:", t)
		case t := <-ticker2.C:
			fmt.Println("Ticker2 at:", t)
		}
	}
}
