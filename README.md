# 新規で参加する場合
`.env`と`serviceAccountKey.json`をもらうこと

## goコンテナだけbuildして再起動する場合
docker-compose up -d --no-deps --build app

## 単体で実行
docker-compose exec app go run main.go

## realizeのログ監視
make attach(alias ath)

## アノテーションコメント
| prefix | 内容 |
| ---- | ---- |
| TODO: |  あとで追加、修正するべき機能がある。 |
| FIXME: | 既知の不具合があるコード。修正が必要。 |
| HACK: | あまりきれいじゃないコード。リファクタリングが必要。 |
| XXX: | 危険！動くけどなぜうごくかわからない。 |
| REVIEW: | 意図した通りに動くか、見直す必要がある。 |
| OPTIMIZE: |  無駄が多く、ボトルネックになっている。 |
| CHANGED: |  コードをどのように変更したか。 |
| NOTE: | なぜ、こうなったという情報を残す。 |
| WARNING: |  注意が必要。 |
