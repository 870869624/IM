createdb:
	docker exec -it wechat-db-1 createdb --username=postgres wechat

migrateup:
	migrate -path db/migrations -database "postgresql://postgres:123456@localhost:5432/wechat?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migrations -database "postgresql://postgres:123456@localhost:5432/wechat?sslmode=disable" -verbose down

sqlc:
	sqlc generate

server:
	go run main.go

.PHONY: createdb migratedown migrateup sqlc server