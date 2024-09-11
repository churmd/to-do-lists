first_time_setup: tools
	go mod download

tools:
	go install github.com/vektra/mockery/v2@v2.45.1

generate:
	go generate -v ./...
	mockery

format:
	go fmt ./...

test: clean_test_cache
	go test ./...

clean_test_cache:
	go clean -testcache

coverage:
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out

update_deps:
	go get -t -u ./...