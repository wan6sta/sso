migrate:
	go run ./cmd/migrator/ --storage-path=./storage/sso.db --migrations-path=./migrations
migrate-tests:
	go run ./cmd/migrator/ --storage-path=./storage/sso.db --migrations-path=./tests/migrations --migrations-table=migrations_tests