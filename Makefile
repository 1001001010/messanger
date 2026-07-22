include apps/server/.env
export

.PHONY: up down logs ps proto lint migrate-up migrate-down server web

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



migrate-up:
	migrate \
		-path apps/server/migrations \
		-database "$(DB_URL)" \
		up

migrate-down:
	migrate \
		-path apps/server/migrations \
		-database "$(DB_URL)" \
		down 1

server:
	cd apps/server && go run ./cmd/server

web:
	cd apps/web && pnpm run dev

sqlc:
	cd apps/server && sqlc generate