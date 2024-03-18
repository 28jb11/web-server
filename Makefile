run:
	@go mod tidy
	@templ generate
	@air
