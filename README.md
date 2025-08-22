# Go言語における小数の検証プロジェクト

Go言語で小数を扱う場合における、いくつかの方法を検証します。

## 検証一覧

- float型を使う
- 整数に直す (万分率)
- Go標準パッケージmath/bigのratを使う
- github.com/shopspring/decimalを使う

## 検証内容

- floatの場合に誤差が生じるケースを検証する
- DBへ保存する
- DBからデータを取得する

## 技術スタック

- 言語: Go (1.25.x)
- データベース: MySQL
- コンテナ: Docker, Docker Compose
