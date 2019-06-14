build:
	docker build -t c4 .
run:
	docker run -i -t --rm c4
test:
	go test -v ./...
test-coverage:
	go test -cover ./...