PG_URL=postgresql://root:secret@localhost:5432/trailer_manager?sslmode=disable

postgres:
	docker run --name postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:14.5-alpine

createdb:
	docker exec -it postgres createdb --username=root --owner=root trailer_manager

dropdb:
	docker exec -it postgres dropdb trailer_manager

dbup:
	migrate -path db/migration -database "$(PG_URL)" -verbose up

dbup1:
	migrate -path db/migration -database "$(PG_URL)" -verbose up 1

dbdown:
	migrate -path db/migration -database "$(PG_URL)" -verbose down

dbdown1:
	migrate -path db/migration -database "$(PG_URL)" -verbose down 1

sqlc:
	sqlc generate

.PHONY: postgres createdb dropdb dbup dbup1 dbdown dbdown1 sqlc