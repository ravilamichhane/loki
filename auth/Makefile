run :
	air run -c .air.toml

migrate :
	@echo "Migrating database..."
	@go run migrations/main.go

gen:
	 protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    authorisation/authorisation.proto