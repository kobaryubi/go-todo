# go-todo

## Getting Started

Run the development server:

```
docker compose -f compose.dev.yml up -d
```

List containers

```
docker compose -f compose.dev.yml ps
```

Stop the development server:

```
docker compose -f compose.dev.yml down
```

Connect to PostgreSQL

```
psql -h localhost -U postgres
```
