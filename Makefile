all: generate

generate:
	go run .github/generate-to-type/to-basic-type.go && gofumpt -l -w . && go test ./internal/. # 基本类型
