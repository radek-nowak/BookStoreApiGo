tidy:
	go mod tidy 

run:
	go run ./cmd/api/main.go

codegen:
	oapi-codegen \
	-generate fiber,types,strict-server,spec \
	-package api -o api/http.gen.go openapi/api.yaml

