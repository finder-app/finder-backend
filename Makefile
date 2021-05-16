attach:
	docker attach finder-backend_app_1

test:
	docker-compose exec app go test -v -cover ./usecase/

test-all:
	docker-compose exec app go test -v -cover ./...

m-file:
	migrate create -ext sql -dir ./db/migrate -seq $(name)

m-up:
	migrate -source file://./db/migrate -database 'mysql://root:finder0501@tcp(localhost:13306)/finder_development' up

m-down:
	migrate -source file://./db/migrate -database 'mysql://root:finder0501@tcp(localhost:13306)/finder_development' down $(n)

m-force:
	migrate -source file://./db/migrate -database 'mysql://root:finder0501@tcp(localhost:13306)/finder_development' force $(v)
