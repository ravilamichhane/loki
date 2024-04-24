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

generator-compile : 
	@echo "Generating CODE GENERATOR"
	@cd generators && go build -o ../generator main.go
	./generator --t="service" --package="coffee" --root="app"   