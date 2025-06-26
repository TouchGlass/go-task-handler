# Makefile для создания миграций

# Переменные которые будут использоваться в наших командах (Таргетах)
DB_DSN := "postgres://postgres:mypassword@localhost:5432/postgres?sslmode=disable"
MIGRATE := migrate -path ./migrations -database $(DB_DSN)

# Таргет для создания новой миграции
migrate-new:
	migrate create -ext sql -dir ./migrations ${NAME}

# Применение миграций
migrate-up:
	$(MIGRATE) up

# Откат миграций
migrate-down:
	$(MIGRATE) down

# Показать текущую версию миграции
migrate-version:
	$(MIGRATE) version

# для удобства добавим команду run
run:
	go run ./cmd/main.go


gen:
	oapi-codegen -config openapi/.openapi-config.yaml -include-tags tasks -package tasks openapi/openapi.yaml > ./internal/web/tasks/api.gen.go