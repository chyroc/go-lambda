all: generate

generate:
	( cd .github/generate-to-type && go build -o /tmp/to-basic-type ./to-basic-type.go )
	( cd .github/generate-to-type && go build -o /tmp/to-basic-type-slice ./to-basic-type-slice.go )
	/tmp/to-basic-type
	/tmp/to-basic-type-slice
	gofumpt -l -w . && go test ./internal/. # 基本类型
