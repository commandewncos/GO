POSTGRESQL_URL="postgres://postgres:secret@localhost:5432/postgres?sslmode=disable"
server:
	@go run command/server/*

dev:
	@go run command/dev/dev.go

up:
	@docker compose up

migrate-create:
	@migrate create -ext sql -dir db/migrations -seq create_notes_table	


migrate-up:
	@migrate -database ${POSTGRESQL_URL} -path db/migrations up

migrate-down:
	@migrate -database ${POSTGRESQL_URL} -path db/migrations down



.PHONY: server dev up 