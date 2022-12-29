postgres:
	docker run --name carContainer -p7777:5432 -e POSTGRES_PASSWORD=1234567891 -d postgres:latest

createDB:
	docker exec -it carContainer createdb --username=postgres --owner=postgres cars

dropDB:
	docker exec -it carContainer dropdb --username=postgres cars

migrateup:
	migrate -path internal/db/migration -database "postgresql://postgres:1234567891@localhost:7777/cars?sslmode=disable" -verbose up

migratedown:
	migrate -path internal/db/migration -database "postgresql://postgres:1234567891@localhost:7777/cars?sslmode=disable" -verbose down

.PHONY: postgres createDB dropDB migrateup migratedown