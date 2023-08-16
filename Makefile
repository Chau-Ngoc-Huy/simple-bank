postgres:
	docker run --name postgres1 -e POSTGRES_PASSWORD=12345678 -e POSTGRES_USER=root -p 5432:5432 -d postgres
createdb:
	docker exec -it postgres1 createdb --username=root --owner=root simple_bank
dropdb:
	docker exec -it postgres1 dropdb simple_bank
migrateup:
	migrate -path db/migration -database "postgresql://root:12345678@localhost:5432/simple_bank?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migration -database "postgresql://root:12345678@localhost:5432/simple_bank?sslmode=disable" -verbose down
sqlc:
	docker run --rm -v "D:\Goland_Project\simple-bank:/src" -w /src kjconroy/sqlc generate