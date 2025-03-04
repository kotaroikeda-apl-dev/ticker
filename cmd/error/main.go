package main

import (
	"fmt"
	"time"
)

func tickerJob() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("エラー発生:", r)
			// エラー発生後、再度 `tickerJob()` を開始
			go tickerJob()
		}
	}()

	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop() // Ticker を確実に停止

	for {
		now := time.Now() // 現在時刻をキャプチャ
		fmt.Println("Tickerジョブ実行:", now)

		select {
		case <-ticker.C:
			// 確実に 10 秒ごとにエラーを発生させる
			if now.Second()%10 == 0 {
				panic("予期しないエラー")
			}
		}
	}
}

func main() {
	go tickerJob()

	time.Sleep(30 * time.Second)
	fmt.Println("メイン処理終了")
}
