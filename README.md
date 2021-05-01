## 新規で参加する場合
`.env`と`serviceAccountKey.json`をもらってください

## db管理
### migration
[`golang-migrate`](https://github.com/golang-migrate/migrate)を使用しています

#### setup
```
brew install golang-migrate
```

#### migration create
~~`name=?`には処理する内容を**スネークケースで**追加してください~~
今使えないからコマンドベタ打ちする

例）
- make migration name=create_users_table
- make migration name=add_column_to_users
```
make migration name=?
migrate create -ext sql -dir ./migration/ddl -seq <処理する内容>
```

#### migrate up
未実行のmigrationが全て実行されます
```
make migrate-up
```

#### migrate down
`n=?`には数値を渡してください
```
make migrate-down n=?
```

## アノテーションコメント
| prefix | 内容 |
| ---- | ---- |
| TODO: |  あとで追加、修正するべき機能がある |
| FIXME: | 既知の不具合があるコード修正が必要 |
| HACK: | あまりきれいじゃないコードリファクタリングが必要 |
| XXX: | 危険！動くけどなぜうごくかわからない |
| REVIEW: | 意図した通りに動くか、見直す必要がある |
| OPTIMIZE: |  無駄が多く、ボトルネックになっている |
| CHANGED: |  コードをどのように変更したか |
| NOTE: | なぜ、こうなったという情報を残す |
| WARNING: |  注意が必要 |

## docker関係
### goコンテナだけbuildして再起動する場合
`docker-compose up -d --no-deps --build app`

### 単体で実行
`docker-compose exec app go run server.go`

### realizeのログ監視
`make attach`(`alias ath`)
