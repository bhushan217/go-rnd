# .PHONY: clean, test, run

DB_MIG=db/migrations
DB_URL=postgresql://root:secret@localhost:5643/simple_blog?sslmode=disable
# clean:
#   rm -rf *.out
.PHONY: build
build:
	go build -o ./bin

migrateup:
	migrate -path "$(DB_MIG)" -database "$(DB_URL)" -verbose up

migrateup1:
	migrate -path "$(DB_MIG)" -database "$(DB_URL)" -verbose up 1

migratedown:
	migrate -path "$(DB_MIG)" -database "$(DB_URL)" -verbose down

migratedown1:
	migrate -path "$(DB_MIG)" -database "$(DB_URL)" -verbose down 1

new_migration:
	migrate create -ext sql -dir "$(DB_MIG)" -seq $(name)


db_docs:
	dbdocs build doc/db.dbml

db_schema:
	dbml2sql --postgres -o doc/schema.sql doc/db.dbml

sqlc:
	sqlc generate

# test:
# 	go test -coverprofile=coverage.out && go tool cover -html=coverage.out
test:
	go test -v -cover -short ./...
run:
	go run main.go
server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/bhushan217/go-rnd/db/sqlc Store
	mockgen -package mockwk -destination worker/mock/distributor.go github.com/bhushan217/go-rnd/worker TaskDistributor

proto:
	rm -f pb/*.go
	rm -f doc/swagger/*.swagger.json
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
	--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
	--openapiv2_out=doc/swagger --openapiv2_opt=allow_merge=true,merge_file_name=simple_bank \
	proto/*.proto
	statik -src=./doc/swagger -dest=./doc

evans:
	evans --host localhost --port 9090 -r repl

redis:
	docker run --name redis -p 6379:6379 -d redis:7-alpine

.PHONY: network postgres createdb dropdb migrateup migratedown migrateup1 migratedown1 new_migration db_docs db_schema sqlc test server mock proto evans redis