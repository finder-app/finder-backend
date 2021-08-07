attach:
	docker attach finder-backend_app_1

test:
	docker-compose exec app go test -v -cover ./interface/controller
# docker-compose exec app go test -v -cover ./usecase
# docker-compose exec app go test -v -cover ./infrastructure/repository

test-all:
	docker-compose exec app go test -v -cover ./...

# NOTE: e.g. make m-file name=create_users
m-file:
	migrate create -ext sql -dir ./db/migrate -seq $(name)

m-up:
	migrate -source file://./db/migrate -database 'mysql://root:finder0501@tcp(localhost:13306)/finder_development' up

# NOTE: n回分migrationが戻るので、実行時は要注意
# NOTE: e.g. make m-down n=7
m-down:
	migrate -source file://./db/migrate -database 'mysql://root:finder0501@tcp(localhost:13306)/finder_development' down $(n)

# NOTE: e.g. make m-force v=7
m-force:
	migrate -source file://./db/migrate -database 'mysql://root:finder0501@tcp(localhost:13306)/finder_development' force $(v)

gqlgen:
	go run github.com/99designs/gqlgen generate