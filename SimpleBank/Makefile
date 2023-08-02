postgres:
	docker run --name postgres -p 5432:5432 -e POSTGRES_PASSWORD=mysecretpassword -d postgres

createdb:
	docker exec -it postgres createdb --username=postgres --owner=postgres simple_bank

dropdb:
	docker exec -it postgres dropdb -U postgres simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://postgres:mysecretpassword@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://postgres:mysecretpassword@localhost:5432/simple_bank?sslmode=disable" -verbose down

sqlc:
	docker run --rm -v "%cd%:/src" -w /src kjconroy/sqlc generate
	
listdb:
	docker exec -it postgres psql -U postgres -l

.PHONY: postgres createdb dropdb migrateup migratedown sqlc