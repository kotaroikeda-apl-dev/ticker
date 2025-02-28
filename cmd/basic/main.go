package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Ticker example started")

	// 1秒ごとに tick を送る Ticker を作成
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop() // 確実に停止

	// 5秒後に停止するためのタイマー
	stopTimer := time.After(5 * time.Second)

	for {
		select {
		case t := <-ticker.C:
			fmt.Println("Tick at:", t)
		case <-stopTimer:
			fmt.Println("Stopping ticker...")
			return
		}
	}
}
