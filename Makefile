test:
	go test -cover ./... -coverprofile=coverage.out && \
	go run github.com/jandelgado/gcov2lcov@v1.0.5 -infile coverage.out -outfile lcov.info

lint:
	golangci-lint run

install-tools:
	go install github.com/jandelgado/gcov2lcov@v1.0.5
