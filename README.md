# GoのTickerを使ったサンプル

## 概要
このプロジェクトは、Goの`time.Ticker`を使って定期的な処理を実行し、そのライフサイクルを適切に管理する方法を示します。また、`select`文を用いたチャンネルの処理や、`return`が関数の実行に与える影響についても学習します。

## `time.Ticker`とは？
`time.Ticker`は、指定した間隔ごとに時間を通知するチャンネル(`ticker.C`)を提供するGoの機能です。定期的な処理を実行する際に便利です。

### 学んだポイント
- **Tickerの作成**: `time.NewTicker(間隔)` で指定した間隔ごとに動作するTickerを作成。
- **Tickの受信**: `ticker.C` から定期的な時間イベントを取得。
- **Tickerの停止**: `ticker.Stop()` を呼び出すことでリソースリークを防ぐ。
- **`defer ticker.Stop()` の活用**: 関数終了時に確実に `ticker.Stop()` を実行。
- **`select` を使ったチャンネル処理**: 複数のチャンネルを同時に監視し、どれかが準備完了すると処理を実行。
- **`return` の影響**: `defer` は関数終了前に必ず実行されることを確認。

## コードサンプル
```go
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Tickerのサンプルを開始")

	// 1秒ごとにTickを送るTickerを作成
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop() // 関数終了時にTickerを停止

	// 5秒後にTickerを停止するタイマー
	stopTimer := time.After(5 * time.Second)

	for {
		select {
		case t := <-ticker.C:
			fmt.Println("Tick at:", t)
		case <-stopTimer:
			fmt.Println("Tickerを停止します...")
			return // `main()` を終了
		}
	}
}
```

## 実行方法
### コードを実行する
```sh
go run ticker_example.go
```

### 期待される出力
```sh
Tickerのサンプルを開始
Tick at: 2025-03-01 12:00:01 +0000 UTC
Tick at: 2025-03-01 12:00:02 +0000 UTC
Tick at: 2025-03-01 12:00:03 +0000 UTC
Tick at: 2025-03-01 12:00:04 +0000 UTC
Tick at: 2025-03-01 12:00:05 +0000 UTC
Tickerを停止します...
```

## 学習ポイント
1. **`defer ticker.Stop()` を使うことでリソースリークを防ぐ**
2. **`select` で複数のチャンネルを監視し、同時に準備完了した場合はランダムに処理される**
3. **`return` すると `defer` が実行されてから関数が終了する**
4. **`ticker.Stop()` を呼ばないと、Tickerが動き続けてメモリリークの原因になる**

## 今後の発展
- 複数の `Ticker` を並行処理で動作させる実験。
- 遅い処理がある場合の `Ticker` の tick の取り扱いを調査。
- `context.Context` を使用してTickerを適切に制御する。

## 作成者
- **池田虎太郎** | [GitHub プロフィール](https://github.com/kotaroikeda-apl-dev)
