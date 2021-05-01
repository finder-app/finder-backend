# 新規で参加する場合
`.env`と`serviceAccountKey.json`をもらうこと

## goコンテナだけbuildして再起動する場合
docker-compose up -d --no-deps --build app

## 単体で実行
docker-compose exec app go run main.go

## realizeのログ監視
make attach(alias ath)
