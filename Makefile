auth-run:
	@echo "Running app Service..."
	@cd auth && make run

auth-migrate:
	@echo "Migrating database..."
	@cd auth && make migrate


tidy:
	@echo "Tidying up go modules..."
	@cd auth && go mod tidy
	@cd nest && go mod tidy

generator-compile : 
	@echo "Generating CODE GENERATOR"
	@cd generators && go build -o ../generator main.go
	./generator --t="service" --package="coffee" --root="app"   

run-service-front:
	@echo "Running Service Front..."
	@cd apps/servicefront && bun dev