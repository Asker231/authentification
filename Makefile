dev:
	go run cmd/main.go
dbstart:
	docker compose up -d
dbstop:
	dicker compose down
migrate:
	go run migrations/main.go		