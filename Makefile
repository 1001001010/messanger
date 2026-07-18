.PHONY: up down logs ps proto lint server web

up:
	docker compose up -d

down:
	docker compose down

logs:
	docker compose logs -f

ps:
	docker compose ps

proto:
	buf generate

lint:
	buf lint

server:
	cd apps/server && go run ./cmd/server

web:
	cd apps/web && pnpm run dev