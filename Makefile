app-run:
	@echo "Running app Service..."
	@cd app && make run

app-migrate:
	@echo "Migrating database..."
	@cd app && make migrate

