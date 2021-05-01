attach:
	docker attach finder-backend_app_1

migrate-up:
	migrate -source file://./db//migrate -database 'mysql://root:finder0501@tcp(localhost:3308)/finder_development' up

migrate-down:
	migrate -source file://./db//migrate -database 'mysql://root:finder0501@tcp(localhost:3308)/finder_development' down $(n)

migration:
	migrate create -ext sql -dir ./db//migrate -seq $(name)