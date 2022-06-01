DB_URL=postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
  go run main.go

mock:
  mockgen -package mockdb -destination db/mock/store.go github.com/NaokiYazawa/simplebank/db/sqlc Store

.PHONY: migrateup migratedown sqlc test server mock