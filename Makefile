protoc:
	protoc --go_out=internal/grpc \
	--go-grpc_out=internal/grpc \
	--go_opt=paths=source_relative \
	--go-grpc_opt=paths=source_relative \
	quiz.proto

server:
	go run main.go server

quiz:
	go run main.go quiz

mocks: #generate mocks for tests
	@rm -R ./internal/grpc/mocks || true
	@mockery --dir ./internal/grpc --output ./internal/grpc/mocks --all
	@rm -R ./internal/core/port/mocks || true
	@mockery --dir ./internal/core/port --output ./internal/core/port/mocks --all

tests: #run unit tests
	@go test -coverprofile=coverage.out ./internal/... -v -json | tee report.json
	@go tool cover -func=coverage.out

tools: #install tools used in this project
	go install github.com/vektra/mockery/v2@v2.17.0
