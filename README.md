# Go の Ticker を使ったサンプル

## 概要

このプロジェクトは、Go の`time.Ticker`を使って定期的な処理を実行し、そのライフサイクルを適切に管理する方法を示します。また、`select`文を用いたチャンネルの処理や、`return`が関数の実行に与える影響についても学習します。

## `time.Ticker`とは？

`time.Ticker`は、指定した間隔ごとに時間を通知するチャンネル(`ticker.C`)を提供する Go の機能です。定期的な処理を実行する際に便利です。

### 学んだポイント

- **Ticker の作成**: `time.NewTicker(間隔)` で指定した間隔ごとに動作する Ticker を作成。
- **Tick の受信**: `ticker.C` から定期的な時間イベントを取得。
- **Ticker の停止**: `ticker.Stop()` を呼び出すことでリソースリークを防ぐ。
- **`defer ticker.Stop()` の活用**: 関数終了時に確実に `ticker.Stop()` を実行。
- **`select` を使ったチャンネル処理**: 複数のチャンネルを同時に監視し、どれかが準備完了すると処理を実行。
- **`return` の影響**: `defer` は関数終了前に必ず実行されることを確認。
- **`context.WithTimeout()` の活用**: 指定時間後に自動で `context` をキャンセル。
- **`context.WithCancel()` の活用**: 外部のイベントで `context` を手動キャンセル。
- **複数の `Ticker` を `context` で制御**: 複数の `Ticker` を一括管理する。
- **`context.WithDeadline()` の活用**: 指定時刻になったら `context` をキャンセル。
- **`recover()` を使ったエラーハンドリング**: `panic` が発生しても `recover()` で回復し、Ticker を再起動する。

## 実行方法

### コードを実行する

```sh
go run cmd/basic/main.go        # 基本の Ticker
go run cmd/with_timeout/main.go  # context.WithTimeout() を使った Ticker
go run cmd/with_cancel/main.go   # context.WithCancel() を使った Ticker
go run cmd/multiple_tickers/main.go  # 複数の Ticker を context で制御
go run cmd/with_deadline/main.go  # context.WithDeadline() を使った Ticker
go run cmd/with_recover/main.go  # recover() を使ったエラーハンドリング付き Ticker
```

## 学習ポイント

1. **`defer ticker.Stop()` を使うことでリソースリークを防ぐ**
2. **`select` で複数のチャンネルを監視し、同時に準備完了した場合はランダムに処理される**
3. **`return` すると `defer` が実行されてから関数が終了する**
4. **`ticker.Stop()` を呼ばないと、Ticker が動き続けてメモリリークの原因になる**
5. **`context.WithTimeout()` を使うと指定時間後に自動で `context` をキャンセルできる**
6. **`context.WithCancel()` は外部イベントによって手動で `context` をキャンセルできる**
7. **`defer cancel()` を忘れずに呼ぶことで、適切に `context` のリソースを解放できる**
8. **複数の `Ticker` を `context` で管理し、一括停止できるようにする**
9. **`context.WithDeadline()` を使い、特定の時刻で `context` をキャンセルする**
10. **`recover()` を使い、`panic` 発生後も `Ticker` を再起動し続ける**

## 作成者

- **池田虎太郎** | [GitHub プロフィール](https://github.com/kotaroikeda-apl-dev)
