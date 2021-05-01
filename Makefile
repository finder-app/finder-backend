attach:
	docker attach finder-backend_app_1

migrate-up:
	migrate -source file://./infrastructure/migration/ddl -database 'mysql://root:finder0501@tcp(localhost:3308)/finder_development' up

migrate-down:
	migrate -source file://./infrastructure/migration/ddl -database 'mysql://root:finder0501@tcp(localhost:3308)/finder_development' down $(count)

migration:
	migrate create -ext sql -dir infrastructure/migration/ddl -seq $(process)