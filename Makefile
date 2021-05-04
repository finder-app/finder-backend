attach:
	docker attach finder-backend_app_1

m-file:
	migrate create -ext sql -dir ./db/migrate -seq $(name)

m-up:
	migrate -source file://./db/migrate -database 'mysql://root:finder0501@tcp(localhost:3308)/finder_development' up

m-down:
	migrate -source file://./db/migrate -database 'mysql://root:finder0501@tcp(localhost:3308)/finder_development' down $(n)

m-force:
	migrate -source file://./db/migrate -database 'mysql://root:finder0501@tcp(localhost:3308)/finder_development' force $(v)
