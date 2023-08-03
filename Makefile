get-buf:
	npm install @bufbuild/buf

generate-spec:
	npx buf generate
	mv swagger/proto/*.swagger.json swagger/
	rm -rf swagger/proto
	mv internal/rpc/pb/proto/*.go internal/rpc/pb/
	rm -rf internal/rpc/pb/proto

migrate-db:
	go build -o ./bin/migrate cmd/migrate/main.go
	DB_PATH=./config/db/postgres.yaml ./bin/migrate

build:
	go build -o ./bin/server cmd/server/main.go

run:
	make build
	CONFIG_PATH=./config/local.yaml DB_PATH=./config/db/postgres.yaml ./bin/server