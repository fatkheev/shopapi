# Имя сервиса из docker-compose.yml
DB_SERVICE = db

# Запуск базы данных
db-up:
	docker-compose up -d $(DB_SERVICE)

# Остановка базы данных
db-down:
	docker-compose down

# Перезапуск базы данных
db-restart: db-down db-up

# Логи базы данных
db-logs:
	docker-compose logs -f $(DB_SERVICE)

# Путь к миграциям
MIGRATIONS_DIR = migrations
DB_URL = postgres://shopuser:shoppass@localhost:5432/shopdb?sslmode=disable

# Применить миграции
migrate-up:
	migrate -path $(MIGRATIONS_DIR) -database "$(DB_URL)" up

# Откатить миграции
migrate-down:
	migrate -path $(MIGRATIONS_DIR) -database "$(DB_URL)" down

# Проверить текущий статус миграций
migrate-version:
	migrate -path $(MIGRATIONS_DIR) -database "$(DB_URL)" version