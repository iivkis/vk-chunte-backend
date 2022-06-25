include .env
export

build-run:
	go build -o server.exe ./cmd/app/main.go && ./server.exe

postgres-start:
	docker start postgres14

postgres-stop:
	docker stop postgres14

postgres-restart:
	docker restart postgres14

migrate-up:
	migrate -path internal/repository/migration -database=${DB_URL} -verbose up

migrate-down:
	migrate -path internal/repository/migration -database=${DB_URL} -verbose down

migrate-down-up:
	migrate -path internal/repository/migration -database=${DB_URL} -verbose down 
	migrate -path internal/repository/migration -database=${DB_URL} -verbose up 

create-db:
	docker exec -it postgres14 createdb --username=root --owner=root vk-chunte

sqlc-gen:
	sqlc generate

.PHONY: postgres-start postgres-stop postgres-restart create-db
.PHONY: migrate-up migrate-down