.PHONY: image
image:
	docker compose build php-go-scylladb

.PHONY: up
up:
	docker compose up -d

.PHONY: sh
sh:
	docker compose exec php-go-scylladb sh

.PHONY: cqlsh
cqsh:
	docker compose exec scylladb cqlsh

.PHONY: down
down:
	docker compose down --remove-orphans
