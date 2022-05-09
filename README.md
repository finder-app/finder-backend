## 使用技術
- Golang 1.15.2
- Docker
- gin(APIサーバー)
- GraphQL(一部)
- Gorm(ORマッパー、一部生sqlも使用)
- gRPC(microservices化予定のため)
- zap(logger)
- MySQL8.0
- firebase(認証認可, 画像upload)
- Clean Archtectureの考え方を取り入れた開発
  - interface層がリクエストを受け取り&返却、usecase層が処理、repository層がdatabaseのCRUD処理を行う

### 機能一覧
今後更新予定
https://docs.google.com/spreadsheets/d/1Y520iO3APF-pI23YjKtTqdSzs1fTduwQ7NaB2hnUnko/edit#gid=0

## db管理
### migration
[`golang-migrate`](https://github.com/golang-migrate/migrate)を使用しています

#### setup
```
brew install golang-migrate
```

コマンドの詳細は`Makefile`を参照

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
`docker-compose exec app go run main.go`

### realizeのログ監視
`make attach`(`alias ath`)

#######################
