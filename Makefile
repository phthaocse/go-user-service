gen-models:
	go/bin/sqlboiler go/bin/sqlboiler-mysql
migrate-up:
	migrate -path db/migration -database "mysql://root:$(PASSWORD)@tcp(localhost:3306)/user_service" -verbose up
migrate-down:
	migrate -path db/migration -database "mysql://root:$(PASSWORD)@tcp(localhost:3306)/user_service" -verbose down
migrate-create:
	migrate create -ext sql -dir db/migration -seq $(NAME)
