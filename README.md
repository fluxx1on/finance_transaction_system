# Finance transaction system
This service is a simple version of virtual payment providing system. It's make simple to transfer money between different accounts.

#### PostgreSQL Migrations

Before migrate check the config/db/postgres.yaml and set your variables.
```
make migrate-db
```

#### Run Server

Run after PosgreSQL db migration.
```
make run
```

Or run docker-compose build:
```
make docker-build
```

### Stack

- gRPC
- PostgreSQL
- RabbitMQ
- JWT Token
- OpenAPI specification
