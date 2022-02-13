postgres:
	docker run --name postgres14 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:14.1-alpine

createdb: 
	docker exec -it postgres14 createdb --username=root --owner=root buildbox

dropdb: 
	docker exec -it postgres14 dropdb buildbox

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/buildbox?sslmode=disable" -verbose up 

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/buildbox?sslmode=disable" -verbose down

test:
	go test -v -cover ./...

sqlc:
	sqlc generate

server:
	air

.PHONY: createdb createdb dropdb migrateup migratedown test sqlc server