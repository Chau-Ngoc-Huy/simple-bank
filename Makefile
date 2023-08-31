postgres:
	docker run --name postgres1 -e POSTGRES_PASSWORD=12345678 -e POSTGRES_USER=root -p 5432:5432 -d postgres
create-db:
	docker exec -it postgres1 createdb --username=root --owner=root simple_bank
drop-db:
	docker exec -it postgres1 dropdb simple_bank
migrate-up:
	migrate -path db/migration -database "postgresql://root:12345678@localhost:5432/simple_bank?sslmode=disable" -verbose up
migrate-down:
	migrate -path db/migration -database "postgresql://root:12345678@localhost:5432/simple_bank?sslmode=disable" -verbose down
migrate-up1:
	migrate -path db/migration -database "postgresql://root:12345678@localhost:5432/simple_bank?sslmode=disable" -verbose up 1
migrate-down1:
	migrate -path db/migration -database "postgresql://root:12345678@localhost:5432/simple_bank?sslmode=disable" -verbose down 1
migrate-create:
	migrate create -ext sql -dir db/migration -seq add_users
sqlc:
	docker run --rm -v "D:\Goland_Project\simple-bank:/src" -w /src kjconroy/sqlc generate
server:
	go run main.go
test:
	go test -v -cover ./...
container:
	docker exec -it postgres1 psql -U root -d simple_bank
mock:
	mockgen -destination db/mock/store.go -package mockdb --build_flags=--mod=mod github.com/simple-bank/db/sqlc Store