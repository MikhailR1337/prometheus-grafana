run:
	go run main.go

simulate:
	go run tests/main.go

up-server:
	docker compose -f docker-compose.yml up -d

down-server:
	docker compose -f docker-compose.yml down
