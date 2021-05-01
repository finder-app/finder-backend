## containerの立ち上げ
docker-compose build

docker-compose up

## goコンテナだけbuildして再起動する場合
docker-compose up -d --no-deps --build app

## 単体で実行
docker-compose exec app go run main.go

## デバッグモード
<!-- 動かない！ -->
docker-compose exec app go-pry run main.go

## realizeのログ監視
make attach(alias ath)

# vscode デバッカー
`.vscode/launch.json`で
`${workspaceFolder}/app`って書いて、
appディレクトリ配下のファイルを読むようにしてる
