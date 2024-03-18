run:
	@go mod tidy
	@templ generate
	@go run .
