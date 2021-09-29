all: generate

generate:
	( cd .github/generate-to-type && go build -o /tmp/to-basic-type ./to-basic-type.go )
	( cd .github/generate-to-type && go build -o /tmp/to-basic-type-slice ./to-basic-type-slice.go )
	( cd .github/generate-to-type && go build -o /tmp/to-basic-type-array ./to-basic-type-array.go )
	/tmp/to-basic-type
	/tmp/to-basic-type-slice
	/tmp/to-basic-type-array
	gofumpt -l -w .
	go test -coverprofile=coverage.txt -coverpkg=./... -parallel 1 -p 1 -count=1 -gcflags=-l ./... && go tool cover -html=coverage.txt
