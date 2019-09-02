APP_NAME=graphql-golang-example
DEV_COMPOSE=deployments/development/docker-compose.yml

# Docker
dev-up:
	docker-compose -f $(DEV_COMPOSE) up -d
dev-down:
	docker-compose -f $(DEV_COMPOSE) down
dev-logs:
	docker-compose -f $(DEV_COMPOSE) logs -f

# DB
migrate-create:
	docker-compose -f $(DEV_COMPOSE) exec server goose -dir=./db/migrations create $(NAME) sql
migrate-up:
	docker-compose -f $(DEV_COMPOSE) exec server go run cmd/$(APP_NAME)/main.go migrate
sqlboiler:
	docker-compose -f $(DEV_COMPOSE) exec server sqlboiler mysql --wipe -o ./internal/rdb -c ./configs/sqlboiler.toml -p rdb --no-auto-timestamps --no-tests

# GraphQL
gqlgen:
	docker-compose -f $(DEV_COMPOSE) exec server go run github.com/99designs/gqlgen -c configs/gqlgen.yml
