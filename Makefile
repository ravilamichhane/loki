app-run:
	@echo "Running app Service..."
	@cd app && make run

app-migrate:
	@echo "Migrating database..."
	@cd app && make migrate


tidy:
	@echo "Tidying up go modules..."
	@cd app && go mod tidy
	@cd nest && go mod tidy