# Finance transaction system
In current time service awailable to provide money between two users with transaction. All transactions stacking on RabbitMQ queues and consuming by workers.

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

### Stack

- gRPC
- PostgreSQL
- RabbitMQ
- JWT Token (in developing)
- OpenAPI specification
