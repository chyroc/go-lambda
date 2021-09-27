all: generate

generate:
	go run .github/generate-to-type/main.go
	gofumpt -l -w .