# go-web-standard

## dhirectory

```
.
├── README.md
├── cmd  - メイン
│   ├── injector.go - DI設定
│   ├── main.go - 起動、終了処理
│   └── wire.go - wire生成コード
├── go.mod
├── go.sum
├── internal - 非公開(cmdからimport)
│   ├── adaptor - アダプタ層
│   │   ├── configuration - 設定関連
│   │   ├── domainimpl - impl実装
│   │   │   ├── repository - repository実装
│   │   │   └── service - ドメインサービス実装
│   │   ├── handler - 外部入力処理
│   │   ├── persistence - sqlboiler生成コード
│   │   ├── presenter - 外部出力処理
│   │   └── registry - コンポーネントレジストリ
│   ├── domain - ドメイン層
│   │   ├── factory - ファクトリ
│   │   ├── model - エンティティ、値オブジェクト
│   │   ├── repository - リポジトリ
│   │   └── service - ドメインサービス
│   └── usecase - ユースケース層
├── pkg - 公開
│   ├── domain - ドメイン層
│   └── usecase - ユースケース層
│   │   └── command - 登録系
│   │   └── document - 参照系
└── proto - protoファイル
└── sqlboiler.toml - sqlboiler設定ファイル
```