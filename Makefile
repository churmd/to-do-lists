first_time_setup:
	go mod download

generate:
	go generate -v ./...

format:
	go fmt ./...

clean_test_cache:
	go clean -testcache

test: clean_test_cache
	go test ./...

coverage:
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out

update_deps:
	go get -t -u ./...