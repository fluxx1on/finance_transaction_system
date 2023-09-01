get-buf:
	npm install @bufbuild/buf

generate-spec:
	python3 scripts/proto_gen.py

delete-db:
	dropdb -h localhost -U lunkli transaction_system

migrate-db:
	go build -o ./bin/migrate cmd/migrate/main.go
	DB_PATH=./config/db/postgres.yaml ./bin/migrate

build:
	go build -o ./bin/server ./cmd/server

run:
	go mod download
	make build
	CONFIG_PATH=./config/local.yaml DB_PATH=./config/db/postgres.yaml ./bin/server

docker-build:
	python3 ./scripts/export_secrets.py
	docker compose up -d

d-down:
	docker compose down